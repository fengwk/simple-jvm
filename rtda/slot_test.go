package rtda

import "testing"

func TestSlot_Integer(t *testing.T) {
	var v1 int32 = 12345
	v2 := SlotToInteger(IntegerToSlot(v1))
	if v1 != v2 {
		t.Error(v1, v2)
	}
}

func TestSlot_Long(t *testing.T) {
	var v1 int64 = 12345
	v2 := SlotToLong(LongToSlot(v1))
	if v1 != v2 {
		t.Error(v1, v2)
	}
}

func TestSlot_Float(t *testing.T) {
	var v1 float32 = 3.14
	v2 := SlotToFloat(FloatToSlot(v1))
	if v1 != v2 {
		t.Error(v1, v2)
	}
}

func TestSlot_Double(t *testing.T) {
	v1 := 3.14
	v2 := SlotToDouble(DoubleToSlot(v1))
	if v1 != v2 {
		t.Error(v1, v2)
	}
}

func TestSlot_Ref(t *testing.T) {
	v1 := &Object{}
	v2 := SlotToRef(RefToSlot(v1))
	if v1 != v2 {
		t.Error(v1, v2)
	}
}
