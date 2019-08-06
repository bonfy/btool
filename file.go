package btool

import (
	"os"
	"path/filepath"
)

// GetFileExt : Get the ext of the file
func GetFileExt(filename string) string {
	return filepath.Ext(filename)
}

// IsFileExist : Return the file exsits or not
func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
