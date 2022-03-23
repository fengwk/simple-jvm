package rtda

// StackFrame 栈帧
type StackFrame struct {
	lower              *StackFrame
	localVariableTable *LocalVariableTable
	operandStack       *OperandStack
}
