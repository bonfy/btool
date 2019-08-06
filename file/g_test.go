package file

import "testing"

func TestGetExt(t *testing.T) {
	in := "./hello/readme.md"
	expected := ".md"
	actual := GetExt(in)
	if actual != expected {
		t.Errorf("GetExt(%s) = %s; expected %s", in, actual, expected)
	}
}
