package yacl

type FileType int

const (
	FileTypeYAML FileType = iota + 1
	FileTypeJSON
	FileTypeGOB
)

type FileParams struct {
	Path     []string
	Filename string
	Type     FileType
}

func FromFile[T any](params *FileParams, config *T) {

}
