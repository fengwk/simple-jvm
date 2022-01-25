package classpath

import "testing"

func TestWildcardEntry_Read(t *testing.T) {
	we, err := NewWildcardEntry("java/*")
	if err != nil {
		t.Error(err)
	}

	_, err = we.Read("Hello.class")
	if err != nil {
		t.Error(err)
	}
}
