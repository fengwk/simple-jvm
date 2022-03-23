package classfile

import (
	"fmt"
	"github.com/fengwk/simple-jvm/common"
	"strings"
)

const (
	METHOD_ACC_PUBLIC       = 0x0001
	METHOD_ACC_PRIVATE      = 0x0002
	METHOD_ACC_PROTECTED    = 0x0004
	METHOD_ACC_STATIC       = 0x0008
	METHOD_ACC_FINAL        = 0x0010
	METHOD_ACC_SYNCHRONIZRD = 0x0020
	METHOD_ACC_BRIDGE       = 0x0040
	METHOD_ACC_VARARGS      = 0x0080
	METHOD_ACC_NATIVE       = 0x0100
	METHOD_ACC_ABSTRACT     = 0x0400
	METHOD_ACC_STRICT       = 0x0800
	METHOD_ACC_SYNTHETIC    = 0x1000
)

type MethodInfo struct {
	AccessFlags     U2
	NameIndex       U2
	DescriptorIndex U2
	AttributesCount U2
	Attributes      []AttributeInfo
}

func (mi *MethodInfo) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%-15s: %x\n", "AccessFlags", mi.AccessFlags))
	sb.WriteString(fmt.Sprintf("%-15s: #%d\n", "NameIndex", mi.NameIndex))
	sb.WriteString(fmt.Sprintf("%-15s: #%d\n", "DescriptorIndex", mi.DescriptorIndex))
	sb.WriteString(fmt.Sprintf("%-15s: %d\n", "AttributesCount", mi.AttributesCount))
	sb.WriteString(fmt.Sprintf("%-15s:", "Attributes"))
	for _, v := range mi.Attributes {
		sb.WriteString("\n")
		sb.WriteString(common.Indent(fmt.Sprintf("%v", v), common.INDENT2))
	}
	return sb.String()
}

func parseMethods(r *ClassReader) (U2, []*MethodInfo, error) {
	methodsCount, err := r.readU2()
	if err != nil {
		return methodsCount, nil, err
	}
	methods := make([]*MethodInfo, 0, int(methodsCount))
	for i := 0; i < int(methodsCount); i++ {
		methodInfo, err := parseMethodInfo(r)
		if err != nil {
			return methodsCount, nil, err
		}
		methods = append(methods, methodInfo)
	}
	return methodsCount, methods, nil
}

func parseMethodInfo(r *ClassReader) (*MethodInfo, error) {
	var mi MethodInfo
	var err error
	mi.AccessFlags, err = r.readU2()
	if err != nil {
		return nil, err
	}
	mi.NameIndex, err = r.readU2()
	if err != nil {
		return nil, err
	}
	mi.DescriptorIndex, err = r.readU2()
	if err != nil {
		return nil, err
	}
	mi.AttributesCount, mi.Attributes, err = parseAttributes(r)
	if err != nil {
		return nil, err
	}
	return &mi, nil
}
