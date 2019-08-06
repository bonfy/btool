package btool

import "testing"

func TestGetFileExt(t *testing.T) {
	in := "./hello/readme.md"
	expected := ".md"
	actual := GetFileExt(in)
	if actual != expected {
		t.Errorf("GetExt(%s) = %s; expected %s", in, actual, expected)
	}
}
