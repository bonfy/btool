package btool

import "testing"

var codeAdmin = 0b000001
var codeBlog = 0b000010 // Writhe Blog
var codeIdea = 0b000100 // Write Idea Linggan

func TestCheckRight(t *testing.T) {
	// Check Blog Right
	goodRight := 2
	if CheckRight(goodRight, codeBlog) != true {
		t.Errorf("CheckRight(%d, %d) = %v; expected true", goodRight, codeBlog, CheckRight(goodRight, codeBlog))
	}

	goodRight2 := 0b011
	if CheckRight(goodRight2, codeBlog) != true {
		t.Errorf("CheckRight(%d, %d) = %v; expected true", goodRight2, codeBlog, CheckRight(goodRight2, codeBlog))
	}

	badRight := 0b101
	if CheckRight(badRight, codeBlog) != false {
		t.Errorf("CheckRight(%d, %d) = %v; expected false", badRight, codeBlog, CheckRight(badRight, codeBlog))
	}

}
