package rtda

import "math"

type Slot struct {
	val uint32
	ref *Object
}

func uint32ToSlot(ui uint32) *Slot {
	return &Slot{val: ui}
}

func uint64ToSlot(ul uint64) (*Slot, *Slot) {
	hi := uint32(ul >> 32)
	lo := uint32(ul)
	return &Slot{val: hi}, &Slot{val: lo}
}

func slotToUint32(s *Slot) uint32 {
	return s.val
}

func slotToUint64(s1 *Slot, s2 *Slot) uint64 {
	return uint64(s1.val)<<32 | uint64(s2.val)
}

func IntegerToSlot(i int32) *Slot {
	return uint32ToSlot(uint32(i))
}

func LongToSlot(l int64) (*Slot, *Slot) {
	return uint64ToSlot(uint64(l))
}

func FloatToSlot(f float32) *Slot {
	return uint32ToSlot(math.Float32bits(f))
}

func DoubleToSlot(d float64) (*Slot, *Slot) {
	return uint64ToSlot(math.Float64bits(d))
}

func RefToSlot(ref *Object) *Slot {
	return &Slot{ref: ref}
}

func SlotToInteger(s *Slot) int32 {
	return int32(slotToUint32(s))
}

func SlotToLong(s1 *Slot, s2 *Slot) int64 {
	return int64(slotToUint64(s1, s2))
}

func SlotToFloat(s *Slot) float32 {
	return math.Float32frombits(slotToUint32(s))
}

func SlotToDouble(s1 *Slot, s2 *Slot) float64 {
	return math.Float64frombits(slotToUint64(s1, s2))
}

func SlotToRef(s *Slot) *Object {
	return s.ref
}
