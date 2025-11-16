package goproject242

import (
	"fmt"
	"os"
)

func GetSize(filePath string) (string, error) {
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return "", err
	}

	if fileInfo.IsDir() {
		dir, err := os.Open(filePath)
		if err != nil {
			return "", err
		}
		defer dir.Close()

		entries, err := dir.Readdir(-1)
		if err != nil {
			return "", err
		}

		var totalSize int64
		for _, entry := range entries {
			if !entry.IsDir() {
				totalSize += entry.Size()
			}
		}
		return fmt.Sprintf("%d %s", totalSize, filePath), nil
	}

	return fmt.Sprintf("%d %s", fileInfo.Size(), filePath), nil
}