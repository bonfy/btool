package btool

import (
	"os"
	"path/filepath"
)

// FileGetExt : Get the ext of the file
func FileGetExt(filename string) string {
	return filepath.Ext(filename)
}

// FileIsExist : Return the file exsits or not
func FileIsExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
