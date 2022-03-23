package classfile

import (
	"fmt"
	"github.com/fengwk/simple-jvm/common"
	"github.com/pkg/errors"
	"strings"
)

const (
	CLASS_ACC_PUBLIC     = 0x0001
	CLASS_ACC_FINAL      = 0x0010
	CLASS_ACC_SUPER      = 0x0020
	CLASS_ACC_INTERFACE  = 0x0200
	CLASS_ACC_ABSTRACT   = 0x0400
	CLASS_ACC_SYNTHETIC  = 0x1000
	CLASS_ACC_ANNOTATION = 0x2000
	CLASS_ACC_ENUM       = 0x4000
)

type U1 uint8
type U2 uint16
type U4 uint32
type U8 uint64

// ClassFile 类文件
type ClassFile struct {
	Magic             U4
	MinorVersion      U2
	MajorVersion      U2
	ConstantPoolCount U2
	ConstantPool      ConstantPool // ConstantPoolCount-1
	AccessFlags       U2
	ThisClass         U2 // 必须是对常量池表中某项的一个有效索引值
	SuperClass        U2 // 要么是0，要么是对常量池表中某项的一个有效索引值
	InterfacesCount   U2
	Interfaces        []U2
	FieldsCount       U2
	Fields            []*FieldInfo
	MethodsCount      U2
	Methods           []*MethodInfo
	AttributesCount   U2
	Attributes        []AttributeInfo
}

func (cf *ClassFile) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%-17s: %x\n", "Magic", cf.Magic))
	sb.WriteString(fmt.Sprintf("%-17s: %d\n", "MinorVersion", cf.MinorVersion))
	sb.WriteString(fmt.Sprintf("%-17s: %d\n", "MajorVersion", cf.MajorVersion))
	sb.WriteString(fmt.Sprintf("%-17s: %x\n", "ConstantPoolCount", cf.ConstantPoolCount))
	sb.WriteString(fmt.Sprintf("%-17s:\n%v\n", "ConstantPool", common.Indent(cf.ConstantPool.String(), common.INDENT2)))
	sb.WriteString(fmt.Sprintf("%-17s: %x\n", "AccessFlags", cf.AccessFlags))
	sb.WriteString(fmt.Sprintf("%-17s: %d\n", "ThisClass", cf.ThisClass))
	sb.WriteString(fmt.Sprintf("%-17s: %d\n", "SuperClass", cf.SuperClass))
	sb.WriteString(fmt.Sprintf("%-17s: %d\n", "InterfacesCount", cf.InterfacesCount))
	sb.WriteString(fmt.Sprintf("%-17s: %v\n", "Interfaces", cf.Interfaces))
	sb.WriteString(fmt.Sprintf("%-17s: %d\n", "FieldsCount", cf.FieldsCount))
	sb.WriteString(fmt.Sprintf("%-17s:", "Fields"))
	for _, v := range cf.Fields {
		sb.WriteString("\n")
		sb.WriteString(common.Indent(fmt.Sprintf("\n%v", v), common.INDENT2))
	}
	sb.WriteString("\n\n")
	sb.WriteString(fmt.Sprintf("%-17s: %d\n", "MethodsCount", cf.MethodsCount))
	sb.WriteString(fmt.Sprintf("%-17s:", "Methods"))
	for _, v := range cf.Methods {
		sb.WriteString("\n")
		sb.WriteString(common.Indent(fmt.Sprintf("\n%v", v), common.INDENT2))
	}
	sb.WriteString("\n\n")
	sb.WriteString(fmt.Sprintf("%-17s: %d\n", "AttributesCount", cf.AttributesCount))
	sb.WriteString(fmt.Sprintf("%-17s:", "Attributes"))
	for _, v := range cf.Attributes {
		sb.WriteString(common.Indent(fmt.Sprintf("\n%v", v), common.INDENT2))
	}
	return sb.String()
}

func Parse(classBytes []byte) (*ClassFile, error) {
	var cf ClassFile
	var err error
	r := NewClassReader(classBytes, &cf)

	cf.Magic, err = parseMagic(r)
	if err != nil {
		return nil, err
	}

	cf.MinorVersion, cf.MajorVersion, err = parseVersion(r)
	if err != nil {
		return nil, err
	}

	cf.ConstantPoolCount, cf.ConstantPool, err = parseConstantPool(r)
	if err != nil {
		return nil, err
	}

	cf.AccessFlags, err = r.readU2()
	if err != nil {
		return nil, err
	}

	cf.ThisClass, err = r.readU2()
	if err != nil {
		return nil, err
	}

	cf.SuperClass, err = r.readU2()
	if err != nil {
		return nil, err
	}

	cf.InterfacesCount, err = r.readU2()
	if err != nil {
		return nil, err
	}

	cf.Interfaces, err = r.readU2s(int(cf.InterfacesCount))
	if err != nil {
		return nil, err
	}

	cf.FieldsCount, cf.Fields, err = parseFields(r)
	if err != nil {
		return nil, err
	}

	cf.MethodsCount, cf.Methods, err = parseMethods(r)
	if err != nil {
		return nil, err
	}

	cf.AttributesCount, cf.Attributes, err = parseAttributes(r)
	if err != nil {
		return nil, err
	}

	return &cf, err
}

func parseMagic(r *ClassReader) (U4, error) {
	magic, err := r.readU4()
	if err != nil {
		return 0, err
	}
	if magic != 0xCAFEBABE {
		return magic, errors.Errorf("Magic error, [Magic=%x]", magic)
	}
	return magic, nil
}

func parseVersion(r *ClassReader) (U2, U2, error) {
	// 主版本号从45开始，当前虚拟机支持到JDK8，支持到52
	minorVersion, err := r.readU2()
	if err != nil {
		return 0, 0, err
	}
	majorVersion, err := r.readU2()
	if err != nil {
		return 0, 0, err
	}
	if (majorVersion == 45) || (majorVersion > 45 && majorVersion <= 52 && minorVersion == 0) {
		return minorVersion, majorVersion, nil
	}
	return minorVersion, majorVersion, errors.Errorf("version error, [MinorVersion=%d, MajorVersion=%d]", minorVersion, majorVersion)
}
