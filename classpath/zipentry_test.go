package classpath

import "testing"

func TestZipEntry_Read(t *testing.T) {
	ze, err := NewZipEntry("java/hello.jar")
	if err != nil {
		t.Error(err)
	}

	_, err = ze.Read("Hello.class")
	if err != nil {
		t.Error(err)
	}
}
