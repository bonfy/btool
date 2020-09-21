package btool

import "testing"

func TestFileGetExt(t *testing.T) {
	in := "./hello/readme.md"
	expected := ".md"
	actual := FileGetExt(in)
	if actual != expected {
		t.Errorf("GetExt(%s) = %s; expected %s", in, actual, expected)
	}
}
