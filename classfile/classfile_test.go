package classfile

import (
	"github.com/fengwk/simple-jvm/classpath"
	"testing"
)

func TestClassFile_Parse(t *testing.T) {
	de, err := classpath.NewDirEntry("java")
	if err != nil {
		t.Error(err)
	}

	classBytes, err := de.Read("test/Test.class")
	if err != nil {
		t.Errorf("error is %v", err)
	}

	_, err = Parse(classBytes)
	if err != nil {
		t.Error(err)
	}
}
