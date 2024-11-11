package utils

import (
	"os"
	"project-advanced-concepts/errors/types"
)

func OpenFile(path string) (*os.File, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, &types.FileNotFoundError{ File: path }
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, 0, &types.FileStatError{ }
	}

	return file, int(fi.Size()), nil
}