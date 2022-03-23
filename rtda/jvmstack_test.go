package rtda

import "testing"

func TestJVMStack(t *testing.T) {
	stack := NewJVMStack()
	f1 := &StackFrame{}
	f2 := &StackFrame{}
	f3 := &StackFrame{}
	stack.Push(f1)
	stack.Push(f2)
	stack.Push(f3)
	if stack.Pop() != f3 {
		t.Error(f3)
	}
	if stack.Pop() != f2 {
		t.Error(f2)
	}
	if stack.Pop() != f1 {
		t.Error(f1)
	}
}
