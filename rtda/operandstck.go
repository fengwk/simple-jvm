package rtda

// OperandStack 操作数栈
type OperandStack struct {
	top       int // 指向下一个元素压入的顶部位置
	maxStacks int
	slots     []*Slot
}

func NewOperandStack(maxStacks int) *OperandStack {
	return &OperandStack{0, maxStacks, make([]*Slot, maxStacks)}
}

func (stack *OperandStack) push(s *Slot) {
	if stack.IsFull() {
		panic("operand stack is full")
	}
	stack.slots[stack.top] = s
	stack.top++
}

func (stack *OperandStack) pop() *Slot {
	if stack.IsEmpty() {
		panic("operand stack is empty")
	}
	stack.top--
	ret := stack.slots[stack.top]
	stack.slots[stack.top] = nil
	return ret
}

func (stack *OperandStack) IsFull() bool {
	return stack.top == stack.maxStacks
}

func (stack *OperandStack) IsEmpty() bool {
	return stack.top == 0
}

func (stack *OperandStack) PushInteger(i int32) {
	stack.push(IntegerToSlot(i))
}

func (stack *OperandStack) PushLong(l int64) {
	hi, lo := LongToSlot(l)
	stack.push(lo)
	stack.push(hi)
}

func (stack *OperandStack) PushFloat(f float32) {
	stack.push(FloatToSlot(f))
}

func (stack *OperandStack) PushDouble(d float64) {
	hi, lo := DoubleToSlot(d)
	stack.push(lo)
	stack.push(hi)
}

func (stack *OperandStack) PushRef(ref *Object) {
	stack.push(RefToSlot(ref))
}

func (stack *OperandStack) PopInteger() int32 {
	return SlotToInteger(stack.pop())
}

func (stack *OperandStack) PopLong() int64 {
	hi, lo := stack.pop(), stack.pop()
	return SlotToLong(hi, lo)
}

func (stack *OperandStack) PopFloat() float32 {
	return SlotToFloat(stack.pop())
}

func (stack *OperandStack) PopDouble() float64 {
	hi, lo := stack.pop(), stack.pop()
	return SlotToDouble(hi, lo)
}

func (stack *OperandStack) PopRef() *Object {
	return SlotToRef(stack.pop())
}
