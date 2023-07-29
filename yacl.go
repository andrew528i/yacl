package yacl

import (
	"github.com/andrew528i/yacl/env"
	"github.com/andrew528i/yacl/file"
	"github.com/andrew528i/yacl/flags"
	"github.com/andrew528i/yacl/utils"
)

type YACL[T any] struct {
	flags *flags.Params
	env   *env.Params
	file  *file.Params

	ignoreFlags bool
}

func New[T any]() *YACL[T] {
	return &YACL[T]{
		flags: flags.DefaultParams(),
		env:   env.DefaultParams(),
		file:  file.DefaultParams(),
	}
}

func (s *YACL[T]) SetEnvPrefix(prefix string) {
	s.env.Prefix = prefix
}

func (s *YACL[T]) SetEnvDelimiter(delimiter string) {
	s.env.Delimiter = delimiter
}

func (s *YACL[T]) AddFilePath(path string) {
	s.file.Paths = append(s.file.Paths, path)
}

func (s *YACL[T]) SetFilename(filename string) {
	s.file.Filename = filename
}

func (s *YACL[T]) SetFlagDelimiter(delimiter string) {
	s.flags.Delimiter = delimiter
}

func (s *YACL[T]) SetFlagFormatFunc(f func([]string)) {
	s.flags.FieldPathFormatFunc = f
}

func (s *YACL[T]) SetIgnoreFlags(v bool) {
	s.ignoreFlags = v
}

func (s *YACL[T]) Parse(defaultConfigs ...*T) (*T, error) {
	var cfg T

	// First use default configs if they provided
	for _, defaultConfig := range defaultConfigs {
		utils.MergeStruct(&cfg, defaultConfig)
	}

	// Merge config values from yaml file
	yamlCfg, err := file.ParseYAML[T](s.file)
	if err != nil {
		if _, ok := err.(*file.NotFound); !ok {
			return nil, err
		}
	} else {
		utils.MergeStruct(&cfg, yamlCfg)
	}

	// Merge config values from json file
	jsonCfg, err := file.ParseJSON[T](s.file)
	if err != nil {
		if _, ok := err.(*file.NotFound); !ok {
			return nil, err
		}
	} else {
		utils.MergeStruct(&cfg, jsonCfg)
	}

	// Merge config values from json file
	binaryCfg, err := file.ParseBinary[T](s.file)
	if err != nil {
		if _, ok := err.(*file.NotFound); !ok {
			return nil, err
		}
	} else {
		utils.MergeStruct(&cfg, binaryCfg)
	}

	// Merge config values from env
	envCfg, err := env.Parse[T](s.env)
	if err != nil {
		return nil, err
	}

	utils.MergeStruct(&cfg, envCfg)

	// Merge config values from command line interface
	if !s.ignoreFlags {
		flagCfg, err := flags.Parse[T](s.flags)
		if err != nil {
			return nil, err
		}

		utils.MergeStruct(&cfg, flagCfg)
	}

	return &cfg, nil
}
