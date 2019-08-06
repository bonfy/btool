package file

import (
	"os"
	"path/filepath"
)

// GetExt : Get the ext of the file
func GetExt(filename string) string {
	return filepath.Ext(filename)
}

// IsExist : Return the file exsits or not
func IsExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
