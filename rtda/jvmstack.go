package rtda

// JVMStack java虚拟机栈
type JVMStack struct {
	//maxSize int
	//size    int
	top *StackFrame
}

func NewJVMStack() *JVMStack {
	return &JVMStack{}
}

func (stack *JVMStack) IsEmpty() bool {
	return stack.top == nil
}

func (stack *JVMStack) Push(frame *StackFrame) {
	//if stack.size == stack.maxSize {
	//	panic("java.lang.StackOverflowError")
	//}
	if stack.top == nil {
		stack.top = frame
	} else {
		frame.lower = stack.top
		stack.top = frame
	}
	//stack.size++
}

func (stack *JVMStack) Pop() *StackFrame {
	if stack.top == nil {
		panic("no frame in stack")
	}
	ret := stack.top
	stack.top = ret.lower
	ret.lower = nil
	//stack.size--
	return ret
}
