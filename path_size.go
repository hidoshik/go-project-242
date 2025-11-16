package goproject242

import (
	"os"
)

func GetSize(filePath string) (int64, error) {
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return 0, err
	}

	if fileInfo.IsDir() {
		dir, err := os.Open(filePath)
		if err != nil {
			return 0, err
		}
		defer dir.Close()

		entries, err := dir.Readdir(-1)
		if err != nil {
			return 0, err
		}

		var totalSize int64
		for _, entry := range entries {
			if !entry.IsDir() {
				totalSize += entry.Size()
			}
		}
		return totalSize, nil
	}
	return fileInfo.Size(), nil
}