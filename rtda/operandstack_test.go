package rtda

import "testing"

func TestOperandStack_IsEmpty(t *testing.T) {
	v1 := int32(123)
	v2 := int64(456)
	v3 := float32(3.14)
	v4 := 998.2
	v5 := &Object{}
	stack := NewOperandStack(7)

	stack.PushInteger(v1)
	stack.PushLong(v2)
	stack.PushFloat(v3)
	stack.PushDouble(v4)
	stack.PushRef(v5)

	if stack.PopRef() != v5 {
		t.Error(v5)
	}
	if stack.PopDouble() != v4 {
		t.Error(v5)
	}
	if stack.PopFloat() != v3 {
		t.Error(v5)
	}
	if stack.PopLong() != v2 {
		t.Error(v5)
	}
	if stack.PopInteger() != v1 {
		t.Error(v5)
	}
}
