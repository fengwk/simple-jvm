package classpath

import "testing"

func TestDirEntry_Read(t *testing.T) {
	de, err := NewDirEntry("java")
	if err != nil {
		t.Error(err)
	}

	_, err = de.Read("Hello.class")
	if err != nil {
		t.Errorf("error is %v", err)
	}
}
