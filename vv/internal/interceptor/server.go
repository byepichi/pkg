package interceptor

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bluekaki/pkg/errors"
	"github.com/bluekaki/pkg/id"
	"github.com/bluekaki/pkg/pbutil"
	"github.com/bluekaki/pkg/vv/internal/protos/gen"
	"github.com/bluekaki/pkg/vv/options"

	protoV1 "github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	// JournalID a random id used by log journal
	JournalID = "journal-id"
	// Authorization used by auth, both gateway and grpc
	Authorization = "authorization"
	// ProxyAuthorization used by signature, both gateway and grpc
	ProxyAuthorization = "proxy-authorization"
	// Date GMT format
	Date = "date"
	// Method http.XXXMethod
	Method = "method"
	// URI url encoded
	URI = "uri"
	// Body string body
	Body = "body"
	// XForwardedFor forwarded for
	XForwardedFor = "x-forwarded-for"
	// XForwardedHost forwarded host
	XForwardedHost = "x-forwarded-host"
)

type NotifyHandler func(alert *errors.AlertMessage)

// MessageValidator generated by bluekaki/pkg/protoc-gen-message-validator
type MessageValidator interface {
	Valid() error
}

// SessionUserinfo mark userinfo in context
type SessionUserinfo struct{}

// SignatureIdentifier mark identifier in context
type SignatureIdentifier struct{}

var toLoggedMetadata = map[string]bool{
	Authorization:      true,
	ProxyAuthorization: true,
	Date:               true,
	Method:             true,
	URI:                true,
	Body:               true,
	XForwardedFor:      true,
	XForwardedHost:     true,
}

var _ Payload = (*restPayload)(nil)
var _ Payload = (*grpcPayload)(nil)

// Payload rest or grpc payload
type Payload interface {
	JournalID() string
	ForwardedByGrpcGateway() bool
	Service() string
	Date() string
	Method() string
	URI() string
	Body() string
	t()
}

type restPayload struct {
	journalID string
	service   string
	date      string
	method    string
	uri       string
	body      string
}

func (r *restPayload) JournalID() string {
	return r.journalID
}

func (r *restPayload) ForwardedByGrpcGateway() bool {
	return true
}

func (r *restPayload) Service() string {
	return r.service
}

func (r *restPayload) Date() string {
	return r.date
}

func (r *restPayload) Method() string {
	return r.method
}

func (r *restPayload) URI() string {
	return r.uri
}

func (r *restPayload) Body() string {
	return r.body
}

func (r *restPayload) t() {}

type grpcPayload struct {
	journalID string
	service   string
	date      string
	method    string
	uri       string
	body      string
}

func (g *grpcPayload) JournalID() string {
	return g.journalID
}

func (g *grpcPayload) ForwardedByGrpcGateway() bool {
	return false
}

func (g *grpcPayload) Service() string {
	return g.service
}

func (g *grpcPayload) Date() string {
	return g.date
}

func (g *grpcPayload) Method() string {
	return g.method
}

func (g *grpcPayload) URI() string {
	return g.uri
}

func (g *grpcPayload) Body() string {
	return g.body
}

func (g *grpcPayload) t() {}

// NewServerInterceptor create a server interceptor
func NewServerInterceptor(logger *zap.Logger, enablePrometheus bool, notify NotifyHandler) *ServerInterceptor {
	return &ServerInterceptor{
		logger:           logger,
		enablePrometheus: enablePrometheus,
		notify:           notify,
	}
}

// ServerInterceptor the server's interceptor
type ServerInterceptor struct {
	logger           *zap.Logger
	enablePrometheus bool
	notify           NotifyHandler
}

// UnaryInterceptor a interceptor for server unary operations
func (s *ServerInterceptor) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ts := time.Now()

	fullMethod := strings.Split(info.FullMethod, "/")
	serviceName := fullMethod[1]
	methodName := fullMethod[2]

	doJournal := false
	if proto.GetExtension(FileDescriptor.Options(info.FullMethod), options.E_Journal).(bool) {
		doJournal = true
	}

	var journalID string
	meta, _ := metadata.FromIncomingContext(ctx)
	if values := meta.Get(JournalID); len(values) == 0 || values[0] == "" {
		journalID = id.JournalID()
		meta.Set(JournalID, journalID)
		ctx = metadata.NewOutgoingContext(ctx, meta)
	} else {
		journalID = values[0]
	}

	defer func() { // double recover for safety
		if p := recover(); p != nil {
			err = errors.Panic(p)
			errVerbose := fmt.Sprintf("got double panic => error: %+v", err)
			if s.notify != nil {
				s.notify((&errors.AlertMessage{
					JournalId:    journalID,
					ErrorVerbose: errVerbose,
				}).Init())
			}

			err = status.New(codes.Internal, "got double panic").Err()
			s.logger.Error(fmt.Sprintf("%s %s", journalID, errVerbose))
		}
	}()

	defer func() {
		grpc.SetHeader(ctx, metadata.Pairs(runtime.MetadataHeaderPrefix+JournalID, journalID))

		if p := recover(); p != nil {
			err = errors.Panic(p)
			errVerbose := fmt.Sprintf("got panic => error: %+v", err)
			if s.notify != nil {
				s.notify((&errors.AlertMessage{
					JournalId:    journalID,
					ErrorVerbose: errVerbose,
				}).Init())
			}

			s, _ := status.New(codes.Internal, "got panic").WithDetails(&pb.Stack{Verbose: errVerbose})
			err = s.Err()
		}

		if err != nil {
			switch err.(type) {
			case errors.BzError:
				bzErr := err.(errors.BzError)
				s, _ := status.New(codes.Code(bzErr.Combcode()), bzErr.Desc()).WithDetails(&pb.Stack{Verbose: fmt.Sprintf("%+v", bzErr.Err())})
				err = s.Err()

			case errors.AlertError:
				alertErr := err.(errors.AlertError)
				if s.notify != nil {
					alert := alertErr.AlertMessage()
					alert.JournalId = journalID

					s.notify(alert)
				}

				bzErr := alertErr.BzError()
				s, _ := status.New(codes.Code(bzErr.Combcode()), bzErr.Desc()).WithDetails(&pb.Stack{Verbose: fmt.Sprintf("%+v", bzErr.Err())})
				err = s.Err()
			}
		}

		if doJournal {
			journal := &pb.Journal{
				Id: journalID,
				Request: &pb.Request{
					Restapi: ForwardedByGrpcGateway(ctx),
					Method:  info.FullMethod,
					Metadata: func() map[string]string {
						meta, _ := metadata.FromIncomingContext(ctx)
						mp := make(map[string]string)
						for key, values := range meta {
							if key == URI {
								mp[key] = QueryUnescape(values[0])
								continue
							}

							if toLoggedMetadata[key] {
								mp[key] = values[0]
							}
						}
						return mp
					}(),
					Payload: func() *anypb.Any {
						if req == nil {
							return nil
						}

						any, _ := anypb.New(req.(proto.Message))
						return any
					}(),
				},
				Response: &pb.Response{
					Code: codes.OK.String(),
					Payload: func() *anypb.Any {
						if resp == nil {
							return nil
						}

						any, _ := anypb.New(resp.(proto.Message))
						return any
					}(),
				},
				Success: err == nil,
			}

			if err != nil {
				s, _ := status.FromError(err)
				journal.Response.Code = s.Code().String()
				journal.Response.Message = s.Message()

				if len(s.Details()) > 0 {
					journal.Response.ErrorVerbose = s.Details()[0].(*pb.Stack).Verbose
				}
				err = status.New(s.Code(), s.Message()).Err() // reset detail
			}

			journal.CostSeconds = time.Since(ts).Seconds()

			if err == nil {
				s.logger.Info("server unary interceptor", zap.Any("journal", marshalJournal(journal)))
			} else {
				s.logger.Error("server unary interceptor", zap.Any("journal", marshalJournal(journal)))
			}
		}

		if s.enablePrometheus {
			method := info.FullMethod

			if http := proto.GetExtension(FileDescriptor.Options(info.FullMethod), annotations.E_Http).(*annotations.HttpRule); http != nil {
				if x, ok := http.GetPattern().(*annotations.HttpRule_Get); ok {
					method = "get " + x.Get
				} else if x, ok := http.GetPattern().(*annotations.HttpRule_Put); ok {
					method = "put " + x.Put
				} else if x, ok := http.GetPattern().(*annotations.HttpRule_Post); ok {
					method = "post " + x.Post
				} else if x, ok := http.GetPattern().(*annotations.HttpRule_Delete); ok {
					method = "delete " + x.Delete
				} else if x, ok := http.GetPattern().(*annotations.HttpRule_Patch); ok {
					method = "patch " + x.Patch
				}
			}

			if alias := proto.GetExtension(FileDescriptor.Options(info.FullMethod), options.E_MetricsAlias).(string); alias != "" {
				method = alias
			}

			if err == nil {
				MetricsRequestCost.WithLabelValues(method).Observe(time.Since(ts).Seconds())
			} else {
				MetricsError.WithLabelValues(method, status.Code(err).String(), err.Error(), journalID).Observe(time.Since(ts).Seconds())
			}
		}
	}()

	if req != nil {
		if validator, ok := req.(MessageValidator); ok {
			if err := validator.Valid(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
	}

	var (
		authorizationValidator      UserinfoHandler
		proxyAuthorizationValidator SignatureHandler
	)

	if option := proto.GetExtension(FileDescriptor.Options(serviceName), options.E_Authorization).(*options.Handler); option != nil {
		authorizationValidator = Validator.AuthorizationValidator(option.Name)
	}
	if option := proto.GetExtension(FileDescriptor.Options(info.FullMethod), options.E_MethodAuthorization).(*options.Handler); option != nil {
		authorizationValidator = Validator.AuthorizationValidator(option.Name)
	}

	if option := proto.GetExtension(FileDescriptor.Options(serviceName), options.E_ProxyAuthorization).(*options.Handler); option != nil {
		proxyAuthorizationValidator = Validator.ProxyAuthorizationValidator(option.Name)
	}
	if option := proto.GetExtension(FileDescriptor.Options(info.FullMethod), options.E_MethodProxyAuthorization).(*options.Handler); option != nil {
		proxyAuthorizationValidator = Validator.ProxyAuthorizationValidator(option.Name)
	}

	if authorizationValidator == nil && proxyAuthorizationValidator == nil {
		return handler(ctx, req)
	}

	var auth, proxyAuth string
	if authHeader := meta.Get(Authorization); len(authHeader) != 0 {
		auth = authHeader[0]
	}
	if proxyAuthHeader := meta.Get(ProxyAuthorization); len(proxyAuthHeader) != 0 {
		proxyAuth = proxyAuthHeader[0]
	}

	var payload Payload
	if forwardedByGrpcGateway(meta) {
		payload = &restPayload{
			journalID: journalID,
			service:   serviceName,
			date:      meta.Get(Date)[0],
			method:    meta.Get(Method)[0],
			uri:       meta.Get(URI)[0],
			body:      meta.Get(Body)[0],
		}

	} else {
		payload = &grpcPayload{
			journalID: journalID,
			service:   serviceName,
			date:      meta.Get(Date)[0],
			method:    methodName,
			uri:       info.FullMethod,
			body: func() string {
				if req == nil {
					return ""
				}

				raw, _ := pbutil.ProtoMessage2JSON(req.(protoV1.Message))
				return string(raw)
			}(),
		}
	}

	if authorizationValidator != nil {
		userinfo, err := authorizationValidator(auth, payload)
		if err != nil {
			s := status.New(codes.Unauthenticated, codes.Unauthenticated.String())
			s, _ = s.WithDetails(&pb.Stack{Verbose: fmt.Sprintf("%+v", err)})
			return nil, s.Err()
		}
		ctx = context.WithValue(ctx, SessionUserinfo{}, userinfo)
	}

	if proxyAuthorizationValidator != nil {
		identifier, ok, err := proxyAuthorizationValidator(proxyAuth, payload)
		if err != nil {
			s := status.New(codes.PermissionDenied, codes.PermissionDenied.String())
			s, _ = s.WithDetails(&pb.Stack{Verbose: fmt.Sprintf("%+v", err)})
			return nil, s.Err()
		}
		if !ok {
			return nil, status.Error(codes.PermissionDenied, "signature does not match")
		}
		ctx = context.WithValue(ctx, SignatureIdentifier{}, identifier)
	}

	return handler(ctx, req)
}

type serverWrappedStream struct {
	grpc.ServerStream
}

func (s *serverWrappedStream) RecvMsg(m interface{}) (err error) {
	return s.ServerStream.RecvMsg(m)
}

func (s *serverWrappedStream) SendMsg(m interface{}) (err error) {
	return s.ServerStream.SendMsg(m)
}

// StreamInterceptor a interceptor for server stream operations
func (s *ServerInterceptor) StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	// TODO
	return errors.New("Not currently supported")

	// return handler(srv, &serverWrappedStream{ServerStream: stream})
}
