package yacl

type YACL[T any] struct {
	defaultConfig *T
	currentConfig *T
}

func New[T any](defaultConfig *T) (*YACL[T], error) {
	yaclInst := &YACL[T]{defaultConfig: defaultConfig}

	if err := yaclInst.resetCurrentConfig(); err != nil {
		return nil, err
	}

	return yaclInst, nil
}

func (s *YACL[T]) Parse() (*T, error) {
	return nil, nil
}

func (s *YACL[T]) resetCurrentConfig() error {
	currentConfig, err := deepCopyStruct(s.defaultConfig)
	if err != nil {
		return err
	}

	s.currentConfig = currentConfig

	return nil
}
