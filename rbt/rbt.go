package rbt

import (
	"bytes"
	"container/list"
	"encoding/json"
	"sync"

	"github.com/bluekaki/pkg/errors"
)

var _ RBTree = (*rbTree)(nil)

type RBTree interface {
	// Add overwrite if id repeated
	Add(val Value) bool
	// Delete by id
	Delete(Value) bool
	// Exists by id
	Exists(Value) bool
	// Search by id
	Search(Value) Value
	Size() uint32
	Empty() bool
	Maximum() []Value
	Minimum() []Value
	PopMaximum() []Value
	PopMinimum() []Value
	Asc() []Value
	Desc() []Value
	String() string
	ToJSON() []byte
}

func New() RBTree {
	return new(rbTree)
}

type rbTree struct {
	sync.RWMutex
	size uint32
	root *node
}

func (t *rbTree) String() string {
	t.RLock()
	defer t.RUnlock()

	buf := bytes.NewBufferString("RedBlackTree\n")
	if t.root != nil {
		output(t.root, "", true, buf)
	}
	return buf.String()
}

func output(root *node, prefix string, isTail bool, buf *bytes.Buffer) {
	if root.R != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}

		output(root.R, newPrefix, false, buf)
	}

	buf.WriteString(prefix)
	if isTail {
		buf.WriteString("└── ")
	} else {
		buf.WriteString("┌── ")
	}

	buf.WriteString(root.String())
	buf.WriteString("\n")

	if root.L != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}

		output(root.L, newPrefix, true, buf)
	}
}

func (t *rbTree) Empty() bool {
	t.RLock()
	defer t.RUnlock()

	return t.size == 0
}

func (t *rbTree) Size() uint32 {
	t.RLock()
	defer t.RUnlock()

	return t.size
}

func (t *rbTree) Asc() []Value {
	t.RLock()
	defer t.RUnlock()

	if t.size == 0 {
		return nil
	}

	values := make([]Value, 0, t.size)

	stack := list.New()
	root := t.root
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.L
		}

		if stack.Len() != 0 {
			v := stack.Back()
			root = v.Value.(*node)
			values = append(values, root.values...) // visit

			root = root.R
			stack.Remove(v)
		}
	}

	return values
}

func (t *rbTree) Desc() []Value {
	t.RLock()
	defer t.RUnlock()

	if t.size == 0 {
		return nil
	}

	values := make([]Value, 0, t.size)

	stack := list.New()
	root := t.root
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.R
		}

		if stack.Len() != 0 {
			v := stack.Back()
			root = v.Value.(*node)
			values = append(values, root.values...) // visit

			root = root.L
			stack.Remove(v)
		}
	}

	return values
}

func (t *rbTree) Maximum() []Value {
	t.RLock()
	defer t.RUnlock()

	if t.root == nil {
		return nil
	}

	root := t.maximum(t.root)
	return root.values
}

func (t *rbTree) PopMaximum() []Value {
	t.Lock()
	defer t.Unlock()

	if t.root == nil {
		return nil
	}

	root := t.maximum(t.root)
	values := make([]Value, len(root.values))
	copy(values, root.values)

	for _, val := range values {
		t.delete(val)
	}
	return values
}

func (t *rbTree) Minimum() []Value {
	t.RLock()
	defer t.RUnlock()

	if t.root == nil {
		return nil
	}

	root := t.minimum(t.root)
	return root.values
}

func (t *rbTree) PopMinimum() []Value {
	t.Lock()
	defer t.Unlock()

	if t.root == nil {
		return nil
	}

	root := t.minimum(t.root)
	values := make([]Value, len(root.values))
	copy(values, root.values)

	for _, val := range values {
		t.delete(val)
	}
	return values
}

func (t *rbTree) Exists(val Value) (ok bool) {
	if val == nil {
		return
	}

	t.RLock()
	defer t.RUnlock()

	x := t.lookup(val)
	if x == nil {
		return
	}

	for _, v := range x.values {
		if ok = v.ID() == val.ID(); ok {
			return
		}
	}
	return
}

func (t *rbTree) Search(val Value) Value {
	if val == nil {
		return nil
	}

	t.RLock()
	defer t.RUnlock()

	x := t.lookup(val)
	if x == nil {
		return nil
	}

	for _, v := range x.values {
		if v.ID() == val.ID() {
			return v
		}
	}
	return nil
}

func (t *rbTree) Marshal() []byte {
	t.Lock()
	defer t.Unlock()

	if t.root == nil {
		return nil
	}

	nodes := make([]*node, 0, t.size)
	curLayer := []*node{t.root}

	for len(curLayer) > 0 {
		var nexLayer []*node

		for _, node := range curLayer {
			nodes = append(nodes, node)
			if node.L != nil {
				nexLayer = append(nexLayer, node.L)
			}
			if node.R != nil {
				nexLayer = append(nexLayer, node.R)
			}
		}

		curLayer = nexLayer
	}

	// TODO marshal nodes to raw
	return nil
}

func (t *rbTree) ToJSON() []byte {
	values := t.Asc()
	slice := make([]json.RawMessage, len(values))
	for i, v := range values {
		slice[i] = json.RawMessage(v.ToJSON())
	}

	raw, _ := json.Marshal(slice)
	return raw
}

func JSON2Tree(raw []byte, json2Value func([]byte) (Value, error)) (RBTree, error) {
	tree := new(rbTree)

	var slice []json.RawMessage
	if err := json.Unmarshal(raw, &slice); err != nil {
		return nil, errors.Wrap(err, "unmarshal tree err")
	}

	for _, raw := range slice {
		val, err := json2Value([]byte(raw))
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal value err")
		}

		tree.Add(val)
	}

	return tree, nil
}
