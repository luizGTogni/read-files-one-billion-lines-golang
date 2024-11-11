package types

import "fmt"

type FileNotFoundError struct {
	File	string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file \"%s\" not found", e.File)
}