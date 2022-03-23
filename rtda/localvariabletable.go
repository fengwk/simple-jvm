package rtda

// LocalVariableTable 局部变量表
type LocalVariableTable struct {
	maxLocals int
	slots     []*Slot
}

func NewLocalVariableTable(maxLocals int) *LocalVariableTable {
	return &LocalVariableTable{maxLocals, make([]*Slot, maxLocals)}
}

func (localVarTab *LocalVariableTable) setInteger(idx int, i int32) {
	if idx < 0 || idx >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	localVarTab.slots[idx] = IntegerToSlot(i)
}

func (localVarTab *LocalVariableTable) getInteger(idx int) int32 {
	if idx < 0 || idx >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	return SlotToInteger(localVarTab.slots[idx])
}

func (localVarTab *LocalVariableTable) setLong(idx int, l int64) {
	if idx < 0 || idx+1 >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	localVarTab.slots[idx], localVarTab.slots[idx+1] = LongToSlot(l)
}

func (localVarTab *LocalVariableTable) getLong(idx int) int64 {
	if idx < 0 || idx+1 >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	return SlotToLong(localVarTab.slots[idx], localVarTab.slots[idx+1])
}

func (localVarTab *LocalVariableTable) setFloat(idx int, f float32) {
	if idx < 0 || idx >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	localVarTab.slots[idx] = FloatToSlot(f)
}

func (localVarTab *LocalVariableTable) getFloat(idx int) float32 {
	if idx < 0 || idx >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	return SlotToFloat(localVarTab.slots[idx])
}

func (localVarTab *LocalVariableTable) setDouble(idx int, d float64) {
	if idx < 0 || idx+1 >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	localVarTab.slots[idx], localVarTab.slots[idx+1] = DoubleToSlot(d)
}

func (localVarTab *LocalVariableTable) getDouble(idx int) float64 {
	if idx < 0 || idx+1 >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	return SlotToDouble(localVarTab.slots[idx], localVarTab.slots[idx+1])
}

func (localVarTab *LocalVariableTable) setRef(idx int, ref *Object) {
	if idx < 0 || idx >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	localVarTab.slots[idx] = RefToSlot(ref)
}

func (localVarTab *LocalVariableTable) getRef(idx int) *Object {
	if idx < 0 || idx >= localVarTab.maxLocals {
		panic("index out of bound")
	}
	return SlotToRef(localVarTab.slots[idx])
}
