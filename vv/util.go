package vv

import (
	"context"
	stderr "errors"

	"github.com/bluekaki/pkg/vv/internal/interceptor"

	"google.golang.org/grpc/metadata"
)

var (
	// ErrNotGrpcContext not a grpc context
	ErrNotGrpcContext = stderr.New("ctx does not contain metadata")

	// ErrNoJournalIDInContext no jouranl_id in ctx
	ErrNoJournalIDInContext = stderr.New("not found journal_id in ctx")
)

// JournalID get journal id from context
func JournalID(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrNotGrpcContext
	}

	id := meta.Get(interceptor.JournalID)
	if len(id) == 0 {
		return "", ErrNoJournalIDInContext
	}

	return id[0], nil
}

// Userinfo get userinfo from context
func Userinfo(ctx context.Context) interface{} {
	if ctx == nil {
		return nil
	}
	return ctx.Value(interceptor.SessionUserinfo{})
}

// SignatureIdentifier get signature identifier from context
func SignatureIdentifier(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	val := ctx.Value(interceptor.SignatureIdentifier{})
	if val == nil {
		return ""
	}

	identifier, ok := val.(string)
	if !ok {
		return ""
	}
	return identifier
}