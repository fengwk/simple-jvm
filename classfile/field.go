package classfile

import (
	"fmt"
	"github.com/fengwk/simple-jvm/common"
	"strings"
)

const (
	FIELD_ACC_PUBLIC    = 0x0001
	FIELD_ACC_PRIVATE   = 0x0002
	FIELD_ACC_PROTECTED = 0x0004
	FIELD_ACC_STATIC    = 0x0008
	FIELD_ACC_FINAL     = 0x0010
	FIELD_ACC_VOLATILE  = 0x0040
	FIELD_ACC_TRANSIENT = 0x0080
	FIELD_ACC_SYNTHETIC = 0x1000
	FIELD_ACC_ENUM      = 0x4000
)

type FieldInfo struct {
	AccessFlags     U2
	NameIndex       U2
	DescriptorIndex U2
	AttributesCount U2
	Attributes      []AttributeInfo
}

func (fi *FieldInfo) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%-15s: %x\n", "AccessFlags", fi.AccessFlags))
	sb.WriteString(fmt.Sprintf("%-15s: #%d\n", "NameIndex", fi.NameIndex))
	sb.WriteString(fmt.Sprintf("%-15s: #%d\n", "DescriptorIndex", fi.DescriptorIndex))
	sb.WriteString(fmt.Sprintf("%-15s: %d\n", "AttributesCount", fi.AttributesCount))
	sb.WriteString(fmt.Sprintf("%-15s:", "Attributes"))
	for _, v := range fi.Attributes {
		sb.WriteString("\n")
		sb.WriteString(common.Indent(fmt.Sprintf("%v", v), common.INDENT2))
	}
	return sb.String()
}

func parseFields(r *ClassReader) (U2, []*FieldInfo, error) {
	fieldsCount, err := r.readU2()
	if err != nil {
		return fieldsCount, nil, err
	}
	c := int(fieldsCount)
	fields := make([]*FieldInfo, 0, c)
	for c > 0 {
		fieldInfo, err := parseFieldInfo(r)
		if err != nil {
			return fieldsCount, nil, err
		}
		fields = append(fields, fieldInfo)
		c--
	}
	return fieldsCount, fields, nil
}

func parseFieldInfo(r *ClassReader) (*FieldInfo, error) {
	var fi FieldInfo
	var err error
	fi.AccessFlags, err = r.readU2()
	if err != nil {
		return nil, err
	}
	fi.NameIndex, err = r.readU2()
	if err != nil {
		return nil, err
	}
	fi.DescriptorIndex, err = r.readU2()
	if err != nil {
		return nil, err
	}
	fi.AttributesCount, fi.Attributes, err = parseAttributes(r)
	if err != nil {
		return nil, err
	}
	return &fi, nil
}
