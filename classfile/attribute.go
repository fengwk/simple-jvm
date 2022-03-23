package classfile

import (
	"fmt"
	"github.com/fengwk/simple-jvm/common"
	"strings"
)

const (
	INNER_CLASS_ACC_PUBLIC     = 0x0001
	INNER_CLASS_ACC_PRIVATE    = 0x0002
	INNER_CLASS_ACC_PROTECTED  = 0x0004
	INNER_CLASS_ACC_STATIC     = 0x0008
	INNER_CLASS_ACC_FINAL      = 0x0010
	INNER_CLASS_ACC_INTERFACE  = 0x0200
	INNER_CLASS_ACC_ABSTRACT   = 0x0400
	INNER_CLASS_ACC_SYNTHETIC  = 0x1000
	INNER_CLASS_ACC_ANNOTATION = 0x2000
	INNER_CLASS_ACC_ENUM       = 0x4000
)

const (
	PARAMETERS_ACC_FINAL    = 0x0010
	PARAMETERS_ACC_ABSTRACT = 0x0400 // 表示该参数并未出现在源文件中，是编译器自动生成的
	PARAMETERS_ACC_MANDATED = 0x8000 // 表示该参数是在源文件中隐式定义的。Java语言中的典型场景是this关键字
)

type AttributeInfo interface {
	AttributeNameIndex() U2
	AttributeLength() U4
}

// BaseAttributeInfo 用于实现AttributeNameIndex方法模板
type BaseAttributeInfo struct {
	attributeNameIndex U2
	attributeLength    U4
}

func (attrInfo *BaseAttributeInfo) AttributeNameIndex() U2 {
	return attrInfo.attributeNameIndex
}

func (attrInfo *BaseAttributeInfo) AttributeLength() U4 {
	return attrInfo.attributeLength
}

type CodeAttribute struct {
	*BaseAttributeInfo
	MaxStack             U2   // 操作数栈最大深度
	MaxLocals            U2   // 局部变量表最大深度
	CodeLength           U4   // 字节码指令长度
	Code                 []U1 // 字节码指令
	ExceptionTableLength U2   // 异常表长度
	ExceptionTable       []*ExceptionTableEntry
	AttributesCount      U2
	Attributes           []AttributeInfo
}

func (code *CodeAttribute) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%v:\n", "Code"))
	sb.WriteString(fmt.Sprintf("%s%-20s: %d\n", common.INDENT2, "MaxStack", code.MaxStack))
	sb.WriteString(fmt.Sprintf("%s%-20s: %d\n", common.INDENT2, "MaxLocals", code.MaxLocals))
	sb.WriteString(fmt.Sprintf("%s%-20s: %d\n", common.INDENT2, "MaxLocals", code.MaxLocals))
	sb.WriteString(fmt.Sprintf("%s%-20s: %d\n", common.INDENT2, "ExceptionTableLength", code.ExceptionTableLength))
	sb.WriteString(fmt.Sprintf("%s%-20s: %d\n", common.INDENT2, "AttributesCount", code.AttributesCount))
	sb.WriteString(fmt.Sprintf("%s%-20s:", common.INDENT2, "Attributes"))
	for _, v := range code.Attributes {
		sb.WriteString("\n")
		sb.WriteString(common.Indent(fmt.Sprintf("%v", v), common.INDENT4))
	}
	return sb.String()
}

type ExceptionTableEntry struct {
	StartPc   U2 // StartPc和EndPc表明了异常处理器在code中的有效范围
	EndPc     U2
	HandlerPc U2 // 异常处理器的起点
	CatchType U2 // 不为0时是对常量池表的一个有效索引，类型为ConstantClassInfo
}

type ExceptionsAttribute struct {
	*BaseAttributeInfo
	NumberOfException   U2
	ExceptionIndexTable []U2 // 常量池标有效引用，ConstantClassInfo
}

func (ex *ExceptionsAttribute) String() string {
	return fmt.Sprintf("%v: #%v", "Exceptions", ex.ExceptionIndexTable)
}

type LineNumberTableAttribute struct {
	*BaseAttributeInfo
	LineNumberTableLength U2
	LineNumberTable       []*LineNumberInfo
}

func (lineTab *LineNumberTableAttribute) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%v:", "LineNumberTable"))
	for _, line := range lineTab.LineNumberTable {
		sb.WriteString("\n")
		sb.WriteString(common.Indent(line.String(), common.INDENT2))
	}
	return sb.String()
}

type LineNumberInfo struct {
	StartPc    U2
	LineNumber U2
}

func (line *LineNumberInfo) String() string {
	return fmt.Sprintf("line%3d: %d", line.LineNumber, line.StartPc)
}

type LocalVariableTableAttribute struct {
	*BaseAttributeInfo
	LocalVariableTableLength U2
	LocalVariableTable       []*LocalVariableInfo
}

func (localVarTab *LocalVariableTableAttribute) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%v:\n", "LocalVariableTable"))
	for _, v := range localVarTab.LocalVariableTable {
		sb.WriteString("\n")
		sb.WriteString(common.Indent(fmt.Sprintf("%v", v), common.INDENT2))
	}
	return sb.String()
}

type LocalVariableInfo struct {
	StartPc         U2
	Length          U2
	NameIndex       U2
	DescriptorIndex U2
	Index           U2
}

func (localVar *LocalVariableInfo) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%-15s: #%d\n", "StartPc", localVar.StartPc))
	sb.WriteString(fmt.Sprintf("%-15s: %d\n", "Length", localVar.Length))
	sb.WriteString(fmt.Sprintf("%-15s: #%d\n", "NameIndex", localVar.NameIndex))
	sb.WriteString(fmt.Sprintf("%-15s: #%d\n", "DescriptorIndex", localVar.DescriptorIndex))
	sb.WriteString(fmt.Sprintf("%-15s: #%d\n", "Index", localVar.Index))
	return sb.String()
}

type LocalVariableTypeTableAttribute struct {
	*BaseAttributeInfo
	LocalVariableTableLength U2
	LocalVariableTypeTable   []*LocalVariableTypeInfo
}

type LocalVariableTypeInfo struct {
	StartPc        U2
	Length         U2
	NameIndex      U2
	SignatureIndex U2
	Index          U2
}

type SourceFileAttribute struct {
	*BaseAttributeInfo
	SourceFileIndex U2
}

func (sf *SourceFileAttribute) String() string {
	return fmt.Sprintf("%s: #%d", "SourceFile", sf.SourceFileIndex)
}

type SourceDebugExtensionAttribute struct {
	*BaseAttributeInfo
	DebugExtension []U1
}

type ConstantValueAttribute struct {
	*BaseAttributeInfo
	ConstantValueIndex U2 // 有效的常量值索引，可以为ConstantInteger、ConstantLong、ConstantFloat、ConstantDouble、ConstantString
}

func (cv *ConstantValueAttribute) String() string {
	return fmt.Sprintf("%v: #%d", "ConstantValue", cv.ConstantValueIndex)
}

type InnerClassesAttribute struct {
	*BaseAttributeInfo
	NumberOfClasses U2
	InnerClasses    []*InnerClassInfo
}

type InnerClassInfo struct {
	InnerClassInfoIndex   U2
	OuterClassInfoIndex   U2
	InnerNameIndex        U2
	InnerClassAccessFlags U2
}

type DeprecatedAttribute struct {
	*BaseAttributeInfo
}

func (deprecated *DeprecatedAttribute) String() string {
	return fmt.Sprintf("%v: true", "Deprecated")
}

type SyntheticAttribute struct {
	*BaseAttributeInfo
}

func (synthetic *SyntheticAttribute) String() string {
	return fmt.Sprintf("%v: true", "Synthetic")
}

// TODO 暂不支持
//type StackMapTableAttribute struct {
//	BaseAttributeInfo
//	NumberOfEntries []*StackMapFrameAttribute
//}
//
//type StackMapFrameAttribute struct {
//}

type SignatureAttribute struct {
	*BaseAttributeInfo
	SignatureIndex U2
}

type BootstrapMethodsAttribute struct {
	*BaseAttributeInfo
	NumBootstrapMethods U2
	BootstrapMethods    []*BootstrapMethod
}

type BootstrapMethod struct {
	BootstrapMethodRef    U2
	NumBootstrapArguments U2
	BootstrapArguments    []U2
}

type MethodParametersAttribute struct {
	*BaseAttributeInfo
	ParametersCount U1
	Parameters      []*Parameters
}

type Parameters struct {
	NameIndex   U2
	AccessFlags U2
}

type UnparsedAttribute struct {
	*BaseAttributeInfo
	Bytes []U1
}

type EnclosingMethod struct {
	*BaseAttributeInfo
	ClassIndex  U2
	MethodIndex U2 // 如果当前类不是直接包含在某个方法中则为0
}

// TODO 暂不支持注解
//type RuntimeVisibleAnnotationsAttribute struct {
//	BaseAttributeInfo
//	NumAnnotations U2
//	Annotations    []Annotation
//}
//
//type Annotation struct {
//	TypeIndex            U2
//	NumElementValuePairs U2
//	ElementValuePairs    []ElementValuePair
//}
//
//type ElementValuePair struct {
//	ElementNameIndex U2
//	value            ElementValue
//}
//
//type ElementValue struct {
//}

func parseAttributes(r *ClassReader) (U2, []AttributeInfo, error) {
	attributesCount, err := r.readU2()
	if err != nil {
		return attributesCount, nil, err
	}
	c := int(attributesCount)
	attributes := make([]AttributeInfo, 0, c)
	for c > 0 {
		attrInfo, err := parseAttributeInfo(r)
		if err != nil {
			return attributesCount, nil, err
		}
		attributes = append(attributes, attrInfo)
		c--
	}
	return attributesCount, attributes, nil
}

func parseAttributeInfo(r *ClassReader) (AttributeInfo, error) {
	attrNameIndex, err := r.readU2()
	if err != nil {
		return nil, err
	}
	attrName, err := r.ClassFile.ConstantPool.GetUtf8(int(attrNameIndex))
	if err != nil {
		return nil, err
	}
	attrLen, err := r.readU4()
	if err != nil {
		return nil, err
	}
	base := &BaseAttributeInfo{attrNameIndex, attrLen}

	switch attrName.AsString() {
	case "Code":
		return parseCodeAttribute(r, base)
	case "ConstantValue":
		return parseConstantValueAttribute(r, base)
	case "Deprecated":
		return parseDeprecatedAttribute(r, base)
	case "Exceptions":
		return parseExceptionsAttribute(r, base)
	case "LineNumberTable":
		return parseLineNumberTableAttribute(r, base)
	case "LocalVariableTable":
		return parseLocalVariableTableAttribute(r, base)
	case "SourceFile":
		return parseSourceFileAttribute(r, base)
	case "Synthetic":
		return parseSyntheticAttribute(r, base)
	default:
		return parseUnparsedAttribute(r, base)
	}
}

func parseCodeAttribute(r *ClassReader, base *BaseAttributeInfo) (*CodeAttribute, error) {
	maxStack, err := r.readU2()
	if err != nil {
		return nil, err
	}
	maxLocals, err := r.readU2()
	if err != nil {
		return nil, err
	}
	codeLength, err := r.readU4()
	if err != nil {
		return nil, err
	}
	code, err := r.readU1s(int(codeLength))
	if err != nil {
		return nil, err
	}
	exceptionTableLength, err := r.readU2()
	if err != nil {
		return nil, err
	}
	exceptionTable := make([]*ExceptionTableEntry, 0, int(exceptionTableLength))
	for i := 0; i < int(exceptionTableLength); i++ {
		entry, err := parseExceptionTableEntry(r)
		if err != nil {
			return nil, err
		}
		exceptionTable = append(exceptionTable, entry)
	}
	attributesCount, attributes, err := parseAttributes(r)
	if err != nil {
		return nil, err
	}
	return &CodeAttribute{
		base,
		maxStack,
		maxLocals,
		codeLength,
		code,
		exceptionTableLength,
		exceptionTable,
		attributesCount,
		attributes,
	}, nil
}

func parseExceptionTableEntry(r *ClassReader) (*ExceptionTableEntry, error) {
	startPc, err := r.readU2()
	if err != nil {
		return nil, err
	}
	endPc, err := r.readU2()
	if err != nil {
		return nil, err
	}
	handlerPc, err := r.readU2()
	if err != nil {
		return nil, err
	}
	catchType, err := r.readU2()
	if err != nil {
		return nil, err
	}
	return &ExceptionTableEntry{startPc, endPc, handlerPc, catchType}, nil
}

func parseConstantValueAttribute(r *ClassReader, base *BaseAttributeInfo) (*ConstantValueAttribute, error) {
	constantValueIndex, err := r.readU2()
	if err != nil {
		return nil, err
	}
	return &ConstantValueAttribute{base, constantValueIndex}, nil
}

func parseDeprecatedAttribute(r *ClassReader, base *BaseAttributeInfo) (*DeprecatedAttribute, error) {
	return &DeprecatedAttribute{base}, nil
}

func parseExceptionsAttribute(r *ClassReader, base *BaseAttributeInfo) (*ExceptionsAttribute, error) {
	numberOfException, err := r.readU2()
	if err != nil {
		return nil, err
	}
	exceptionIndexTable, err := r.readU2s(int(numberOfException))
	if err != nil {
		return nil, err
	}
	return &ExceptionsAttribute{base, numberOfException, exceptionIndexTable}, nil
}

func parseLineNumberTableAttribute(r *ClassReader, base *BaseAttributeInfo) (*LineNumberTableAttribute, error) {
	lineNumberTableLen, err := r.readU2()
	if err != nil {
		return nil, err
	}
	lineNumberTable := make([]*LineNumberInfo, 0, int(lineNumberTableLen))
	for i := 0; i < int(lineNumberTableLen); i++ {
		lineNumberInfoAttr, err := parseLineNumberInfo(r)
		if err != nil {
			return nil, err
		}
		lineNumberTable = append(lineNumberTable, lineNumberInfoAttr)
	}
	return &LineNumberTableAttribute{base, lineNumberTableLen, lineNumberTable}, nil
}

func parseLineNumberInfo(r *ClassReader) (*LineNumberInfo, error) {
	startPc, err := r.readU2()
	if err != nil {
		return nil, err
	}
	lineNumber, err := r.readU2()
	if err != nil {
		return nil, err
	}
	return &LineNumberInfo{startPc, lineNumber}, nil
}

func parseLocalVariableTableAttribute(r *ClassReader, base *BaseAttributeInfo) (*LocalVariableTableAttribute, error) {
	localVariableTableLen, err := r.readU2()
	if err != nil {
		return nil, err
	}
	localVariableTable := make([]*LocalVariableInfo, 0, int(localVariableTableLen))
	for i := 0; i < int(localVariableTableLen); i++ {
		localVariableInfoAttr, err := parseLocalVariableInfo(r)
		if err != nil {
			return nil, err
		}
		localVariableTable = append(localVariableTable, localVariableInfoAttr)
	}
	return &LocalVariableTableAttribute{base, localVariableTableLen, localVariableTable}, nil
}

func parseLocalVariableInfo(r *ClassReader) (*LocalVariableInfo, error) {
	startPc, err := r.readU2()
	if err != nil {
		return nil, err
	}
	length, err := r.readU2()
	if err != nil {
		return nil, err
	}
	nameIndex, err := r.readU2()
	if err != nil {
		return nil, err
	}
	descriptorIndex, err := r.readU2()
	if err != nil {
		return nil, err
	}
	index, err := r.readU2()
	if err != nil {
		return nil, err
	}
	return &LocalVariableInfo{startPc, length, nameIndex, descriptorIndex, index}, nil
}

func parseSourceFileAttribute(r *ClassReader, base *BaseAttributeInfo) (*SourceFileAttribute, error) {
	sourceFileIndex, err := r.readU2()
	if err != nil {
		return nil, err
	}
	return &SourceFileAttribute{base, sourceFileIndex}, nil
}

func parseSyntheticAttribute(r *ClassReader, base *BaseAttributeInfo) (*SyntheticAttribute, error) {
	return &SyntheticAttribute{base}, nil
}

func parseUnparsedAttribute(r *ClassReader, base *BaseAttributeInfo) (*UnparsedAttribute, error) {
	bs, err := r.readU1s(int(base.attributeNameIndex))
	if err != nil {
		return nil, err
	}
	return &UnparsedAttribute{base, bs}, nil
}
