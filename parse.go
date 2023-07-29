package yacl

import (
	"github.com/andrew528i/yacl/env"
	"github.com/andrew528i/yacl/file"
	"github.com/andrew528i/yacl/flags"
	"github.com/andrew528i/yacl/utils"
)

func Parse[T any](defaultConfigs ...*T) (*T, error) {
	var cfg T

	// First use default configs if they provided
	for _, defaultConfig := range defaultConfigs {
		utils.MergeStruct(&cfg, defaultConfig)
	}

	// Merge config values from yaml file
	yamlCfg, err := file.ParseYAML[T](file.DefaultParams())
	if err != nil {
		if _, ok := err.(*file.NotFound); !ok {
			return nil, err
		}
	} else {
		utils.MergeStruct(&cfg, yamlCfg)
	}

	// Merge config values from json file
	jsonCfg, err := file.ParseJSON[T](file.DefaultParams())
	if err != nil {
		if _, ok := err.(*file.NotFound); !ok {
			return nil, err
		}
	} else {
		utils.MergeStruct(&cfg, jsonCfg)
	}

	// Merge config values from json file
	binaryCfg, err := file.ParseBinary[T](file.DefaultParams())
	if err != nil {
		if _, ok := err.(*file.NotFound); !ok {
			return nil, err
		}
	} else {
		utils.MergeStruct(&cfg, binaryCfg)
	}

	// Merge config values from env
	envCfg, err := env.Parse[T](env.DefaultParams())
	if err != nil {
		return nil, err
	}

	utils.MergeStruct(&cfg, envCfg)

	// Merge config values from command line interface
	flagCfg, err := flags.Parse[T](flags.DefaultParams())
	if err != nil {
		return nil, err
	}

	utils.MergeStruct(&cfg, flagCfg)

	return &cfg, nil
}
