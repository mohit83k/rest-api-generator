package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

//FileExists check if file exist in given directory in config or specified path
func FileExists(defaultDir string, filename string) (bool, string) {
	file := filepath.Join(defaultDir, filename)

	var exists bool
	exists = fileExists(file)
	if !exists {
		fmt.Println("Checking specified path : ", filename)
		exists = fileExists(filename)
		file = filename
	}
	return exists, file
}

func fileExists(filename string) bool {
	fs, err := os.Stat(filename)

	return err == nil && !fs.IsDir()
}
