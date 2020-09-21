package btool

import "testing"

var codeAdmin = 0b000001
var codeBlog = 0b000010 // Writhe Blog
var codeIdea = 0b000100 // Write Idea Linggan

func TestAuthCheckRight(t *testing.T) {
	// Check Blog Right
	goodRight := 2
	if AuthCheckRight(goodRight, codeBlog) != true {
		t.Errorf("AuthCheckRight(%d, %d) = %v; expected true", goodRight, codeBlog, AuthCheckRight(goodRight, codeBlog))
	}

	goodRight2 := 0b011
	if AuthCheckRight(goodRight2, codeBlog) != true {
		t.Errorf("AuthCheckRight(%d, %d) = %v; expected true", goodRight2, codeBlog, AuthCheckRight(goodRight2, codeBlog))
	}

	badRight := 0b101
	if AuthCheckRight(badRight, codeBlog) != false {
		t.Errorf("AuthCheckRight(%d, %d) = %v; expected false", badRight, codeBlog, AuthCheckRight(badRight, codeBlog))
	}

}
