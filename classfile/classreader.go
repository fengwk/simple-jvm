package classfile

import (
	"encoding/binary"
	"io"
)

type Reader []byte

func NewReader(classBytes []byte) Reader {
	return classBytes
}

func (r *Reader) readU1() (U1, error) {
	if len(*r) < 1 {
		return 0, io.EOF
	}
	v := U1((*r)[0])
	*r = (*r)[1:]
	return v, nil
}

func (r *Reader) readU2() (U2, error) {
	if len(*r) < 2 {
		return 0, io.EOF
	}
	v := U2(binary.BigEndian.Uint16(*r))
	*r = (*r)[2:]
	return v, nil
}

func (r *Reader) readU4() (U4, error) {
	if len(*r) < 4 {
		return 0, io.EOF
	}
	v := U4(binary.BigEndian.Uint32(*r))
	*r = (*r)[4:]
	return v, nil
}

func (r *Reader) readU8() (U8, error) {
	if len(*r) < 8 {
		return 0, io.EOF
	}
	v := U8(binary.BigEndian.Uint64(*r))
	*r = (*r)[8:]
	return v, nil
}

func (r *Reader) readU1s(len int) ([]U1, error) {
	u1s := make([]U1, 0, len)
	for i := 0; i < len; i++ {
		u1, err := r.readU1()
		if err != nil {
			return nil, err
		}
		u1s = append(u1s, u1)
	}
	return u1s, nil
}

func (r *Reader) readU2s(len int) ([]U2, error) {
	u2s := make([]U2, 0, len)
	for i := 0; i < len; i++ {
		u2, err := r.readU2()
		if err != nil {
			return nil, err
		}
		u2s = append(u2s, u2)
	}
	return u2s, nil
}

type ClassReader struct {
	Reader
	ClassFile *ClassFile
}

func NewClassReader(classBytes []byte, cf *ClassFile) *ClassReader {
	r := ClassReader{NewReader(classBytes), cf}
	return &r
}
