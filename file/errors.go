package file

import "fmt"

type NotFound struct {
	Filename string
	Paths    []string
}

func NewNotFound(filename string, paths []string) *NotFound {
	return &NotFound{
		Filename: filename,
		Paths:    paths,
	}
}

func (s NotFound) Error() string {
	return fmt.Sprintf("%v not found in paths: %v", s.Filename, s.Paths)
}
