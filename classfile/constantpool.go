package classfile

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"strings"
)

const (
	CONSTANT_Utf8               = 1
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_Class              = 7
	CONSTANT_String             = 8
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_NameAndType        = 12
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantPool []ConstantInfo // ConstantPoolCount-1

func (cp ConstantPool) String() string {
	sb := strings.Builder{}
	for i, v := range cp {
		if v != nil {
			sb.WriteString(fmt.Sprintf("#%-3d = %v", i, v))
			if i < len(cp)-1 {
				sb.WriteString("\n")
			}
		}
	}
	return sb.String()
}

func (cp ConstantPool) Get(i int) (ConstantInfo, error) {
	if i == 0 {
		return nil, nil
	}
	if i < 0 || i >= len(cp) {
		return nil, errors.Errorf("index out of constant pool bound, [i=%d]", i)
	}
	return cp[i], nil
}

func (cp ConstantPool) GetUtf8(i int) (*ConstantUtf8Info, error) {
	info, err := cp.Get(i)
	if err != nil {
		return nil, err
	}
	if utf8Info, ok := info.(*ConstantUtf8Info); ok {
		return utf8Info, nil
	}
	return nil, errors.Errorf("constant type error, [i=%d, type=%T]", i, info)
}

type ConstantInfo interface {
	Tag() U1
}

type BaseConstantInfo struct {
	tag U1
}

func (consInfo BaseConstantInfo) Tag() U1 {
	return consInfo.tag
}

type ConstantUtf8Info struct {
	BaseConstantInfo
	Length U2
	Bytes  []U1
}

func (utf8Info *ConstantUtf8Info) AsString() string {
	bs := make([]byte, 0, len(utf8Info.Bytes))
	for _, u1 := range utf8Info.Bytes {
		bs = append(bs, byte(u1))
	}
	return string(bs)
}

func (utf8Info *ConstantUtf8Info) String() string {
	return fmt.Sprintf("%-18s %s", "Utf8", utf8Info.AsString())
}

type ConstantIntegerInfo struct {
	BaseConstantInfo
	Bytes U4 // 大端序存储的int
}

func (intInfo *ConstantIntegerInfo) AsInt32() int32 {
	return int32(intInfo.Bytes)
}

func (intInfo *ConstantIntegerInfo) String() string {
	return fmt.Sprintf("%-18s %d", "Integer", intInfo.AsInt32())
}

// 正无穷 Ox7f800000
// 负无穷 Oxff800000
// NaN   Ox7f800001～Ox7fffffff或者Oxff800001～Oxffffffff
// float 算法
// int s = ((bits >> 31) == 0) ? 1 : -1;
// int e = ((bits >> 23) & Oxff);
// int m = (e == O) ?
//     (bits & 0x7fffff) << 1 :
//     (bits & 0x7fffff) | 0x800000;
// s * m * power(2, e-150)
type ConstantFloatInfo struct {
	BaseConstantInfo
	Bytes U4 // 大端序存储，IEEE754单精度浮点格式
}

func (floatInfo *ConstantFloatInfo) AsFloat() float32 {
	return math.Float32frombits(uint32(floatInfo.Bytes))

	// TODO bug
	//var s, e, m float64
	//if (b >> 31) == 0 {
	//	s = 1
	//} else {
	//	s = -1
	//}
	//e = float64((b >> 23) & 0xff)
	//if e == 0 {
	//	m = float64((b & 0x7fffff) << 1)
	//} else {
	//	m = float64((b & 0x7fffff) | 0x800000)
	//}
	//return float32(s * m * math.Pow(2, e-150))
}

func (floatInfo *ConstantFloatInfo) String() string {
	return fmt.Sprintf("%-18s %f", "Float", floatInfo.AsFloat())
}

type ConstantLongInfo struct {
	BaseConstantInfo
	HighBytes U4 // 大端，((long) high_bytes << 32) + low_bytes
	LowBytes  U4 // 大端，((long) high_bytes << 32) + low_bytes
}

func (longInfo *ConstantLongInfo) AsInt64() int64 {
	var hi = uint64(longInfo.HighBytes)
	var lo = uint64(longInfo.LowBytes)
	return int64((hi << 32) + lo)
}

func (longInfo *ConstantLongInfo) String() string {
	return fmt.Sprintf("%-18s %d", "Long", longInfo.AsInt64())
}

// 正无穷 Ox7ff0000000000000L
// 负无穷 OxfffOOOOOOOOOOOOOL
// NaN   Ox7ff0000000000001L～Ox7fffffffffffffffL || OxfffOOOOOOOOOOOOlL～OxffffffffffffffffL
// double 算法
// int s = ((bits >> 63) == 0) ? 1 : -1;
// int e = (int) ((bits>52)& 0x7ffL);
// long m = (e == O) ?
//     (bits & 0xfffffffffffffL) << 1 :
//     (bits & 0xfffffffffffffL) | 0x00000000000000L;
// s * m * power(2, e-1075)
type ConstantDoubleInfo struct {
	BaseConstantInfo
	HighBytes U4 // 大端，((long) high_bytes << 32) + low_bytes
	LowBytes  U4
}

func (doubleInfo *ConstantDoubleInfo) AsDouble() float64 {
	var hi = uint64(doubleInfo.HighBytes)
	var lo = uint64(doubleInfo.LowBytes)
	b := (hi << 32) + lo
	return math.Float64frombits(b)

	// TODO bug
	//var s, e, m float64
	//if (b >> 63) == 0 {
	//	s = 1
	//} else {
	//	s = -1
	//}
	//e = float64((b >> 52) & 0x7ff)
	//if e == 0 {
	//	m = float64((b & 0xfffffffffffff) << 1)
	//} else {
	//	m = float64((b & 0xfffffffffffff) | 0x00000000000000)
	//}
	//return s * m * math.Pow(2, e-1075)
}

func (doubleInfo *ConstantDoubleInfo) String() string {
	return fmt.Sprintf("%-18s %f", "Double", doubleInfo.AsDouble())
}

type ConstantClassInfo struct {
	BaseConstantInfo
	NameIndex U2 // 必须是有效索引，指向ConstantUtf8Info，［Ljava/lang/Thread；
}

func (classInfo *ConstantClassInfo) String() string {
	return fmt.Sprintf("%-18s #%d", "Class", classInfo.NameIndex)
}

type ConstantStringInfo struct {
	BaseConstantInfo
	Index U2
}

func (stringInfo *ConstantStringInfo) String() string {
	return fmt.Sprintf("%-18s #%d", "String", stringInfo.Index)
}

type ConstantFieldRefInfo struct {
	BaseConstantInfo
	ClassIndex       U2 // 必须是有效索引，指向ConstantClassInfo
	NameAndTypeIndex U2 // 必须是有效索引，指向ConstantNameAndTypeInfo
}

func (fieldRefInfo *ConstantFieldRefInfo) String() string {
	return fmt.Sprintf("%-18s #%d.#%d", "Fieldref", fieldRefInfo.ClassIndex, fieldRefInfo.NameAndTypeIndex)
}

type ConstantMethodRefInfo struct {
	BaseConstantInfo
	ClassIndex       U2 // 必须是有效索引，指向ConstantClassInfo
	NameAndTypeIndex U2 // 必须是有效索引，指向ConstantNameAndTypeInfo
}

func (methodRefInfo *ConstantMethodRefInfo) String() string {
	return fmt.Sprintf("%-18s #%d.#%d", "Methodref", methodRefInfo.ClassIndex, methodRefInfo.NameAndTypeIndex)
}

type ConstantInterfaceMethodRefInfo struct {
	BaseConstantInfo
	ClassIndex       U2 // 必须是有效索引，指向ConstantClassInfo
	NameAndTypeIndex U2 // 必须是有效索引，指向ConstantNameAndTypeInfo
}

func (interfaceMethodRefInfo *ConstantInterfaceMethodRefInfo) String() string {
	return fmt.Sprintf("%-18s #%d.#%d", "InterfaceMethodref", interfaceMethodRefInfo.ClassIndex, interfaceMethodRefInfo.NameAndTypeIndex)
}

type ConstantNameAndTypeInfo struct {
	BaseConstantInfo
	NameIndex       U2 // 必须是有效索引，指向ConstantUtf8Info
	DescriptorIndex U2 // 必须是有效索引，指向ConstantUtf8Info
}

func (nameAndTypeInfo *ConstantNameAndTypeInfo) String() string {
	return fmt.Sprintf("%-18s #%d:#%d", "NameAndType", nameAndTypeInfo.NameIndex, nameAndTypeInfo.DescriptorIndex)
}

type ConstantMethodHandleInfo struct {
	BaseConstantInfo
	ReferenceKind  U1 // 1~9
	ReferenceIndex U2
}

func (methodHandleInfo *ConstantMethodHandleInfo) String() string {
	return fmt.Sprintf("%-18s %d,#%d", "MethodHandle", methodHandleInfo.ReferenceKind, methodHandleInfo.ReferenceIndex)
}

type ConstantMethodtypeInfo struct {
	BaseConstantInfo
	DescriptorIndex U2 // 必须是有效索引，指向ConstantUtf8Info
}

func (methodtypeInfo *ConstantMethodtypeInfo) String() string {
	return fmt.Sprintf("%-18s #%d", "Methodtype", methodtypeInfo.DescriptorIndex)
}

type ConstantInvokeDynamicInfo struct {
	BaseConstantInfo
	BootstrapMethodAttrIndex U2 // 必须有效索引，指向属性表，指向BootstrapMethodsAttribute
	NameAndTypeIndex         U2 // 必须是有效索引，指向ConstantNameAndTypeInfo
}

func (invokeDynamicInfo *ConstantInvokeDynamicInfo) String() string {
	return fmt.Sprintf("%-18s $%d,#%d", "InvokeDynamic", invokeDynamicInfo.BootstrapMethodAttrIndex, invokeDynamicInfo.NameAndTypeIndex)
}

func parseConstantPool(r *ClassReader) (U2, ConstantPool, error) {
	constantPoolCount, err := r.readU2()
	if err != nil {
		return 0, nil, err
	}
	constantPool := make([]ConstantInfo, int(constantPoolCount))
	for i := 1; i < int(constantPoolCount); {
		cpInfo, step, err := parseCpInfo(r)
		if err != nil {
			fmt.Println(ConstantPool(constantPool).String())
			return constantPoolCount, nil, err
		}
		constantPool[i] = cpInfo
		i += step
	}
	return constantPoolCount, constantPool, nil
}

func parseCpInfo(r *ClassReader) (ConstantInfo, int, error) {
	t, err := r.readU1()
	if err != nil {
		return nil, 0, err
	}
	switch t {

	case CONSTANT_Utf8:
		l, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		bs, err := r.readU1s(int(l))
		if err != nil {
			return nil, 0, err
		}
		return &ConstantUtf8Info{BaseConstantInfo{t}, l, bs}, 1, nil

	case CONSTANT_Integer:
		bs, err := r.readU4()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantIntegerInfo{BaseConstantInfo{t}, bs}, 1, nil

	case CONSTANT_Float:
		bs, err := r.readU4()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantFloatInfo{BaseConstantInfo{t}, bs}, 1, nil

	case CONSTANT_Long:
		hi, err := r.readU4()
		if err != nil {
			return nil, 0, err
		}
		lo, err := r.readU4()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantLongInfo{BaseConstantInfo{t}, hi, lo}, 2, nil

	case CONSTANT_Double:
		hi, err := r.readU4()
		if err != nil {
			return nil, 0, err
		}
		lo, err := r.readU4()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantDoubleInfo{BaseConstantInfo{t}, hi, lo}, 2, nil

	case CONSTANT_Class:
		i, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantClassInfo{BaseConstantInfo{t}, i}, 1, nil

	case CONSTANT_String:
		i, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantStringInfo{BaseConstantInfo{t}, i}, 1, nil

	case CONSTANT_Fieldref:
		i1, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		i2, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantFieldRefInfo{BaseConstantInfo{t}, i1, i2}, 1, nil

	case CONSTANT_Methodref:
		i1, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		i2, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantMethodRefInfo{BaseConstantInfo{t}, i1, i2}, 1, nil

	case CONSTANT_InterfaceMethodref:
		i1, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		i2, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantInterfaceMethodRefInfo{BaseConstantInfo{t}, i1, i2}, 1, nil

	case CONSTANT_NameAndType:
		i1, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		i2, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantNameAndTypeInfo{BaseConstantInfo{t}, i1, i2}, 1, nil

	case CONSTANT_MethodHandle:
		k, err := r.readU1()
		if err != nil {
			return nil, 0, err
		}
		i, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantMethodHandleInfo{BaseConstantInfo{t}, k, i}, 1, nil

	case CONSTANT_MethodType:
		i, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantMethodtypeInfo{BaseConstantInfo{t}, i}, 1, nil

	case CONSTANT_InvokeDynamic:
		i1, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		i2, err := r.readU2()
		if err != nil {
			return nil, 0, err
		}
		return &ConstantInvokeDynamicInfo{BaseConstantInfo{t}, i1, i2}, 1, nil

	default:
		return nil, 0, errors.Errorf("constant BaseConstantInfo error, [BaseConstantInfo=%d]", t)
	}
}
