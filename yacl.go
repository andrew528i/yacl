package yacl

type YACL[T any] struct {
	defaultConfig *T
	currentConfig *T
}

func New[T any](defaultConfig *T) (*YACL[T], error) {
	yaclInst := &YACL[T]{defaultConfig: defaultConfig}

	return yaclInst, nil
}

func (s *YACL[T]) Parse() (*T, error) {
	return nil, nil
}
