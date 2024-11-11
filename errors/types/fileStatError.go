package types

type FileStatError struct {
}

func (e *FileStatError) Error() string {
	return "file stat failed"
}