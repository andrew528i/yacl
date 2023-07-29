package file

import (
	"os"
)

var DefaultFilename = "config"
var DefaultPath = ""

type Params struct {
	Paths    []string
	Filename string
}

func DefaultParams(extraPaths ...string) *Params {
	paths := make([]string, 0)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	paths = append(paths, cwd)
	paths = append(paths, extraPaths...)

	if DefaultPath != "" {
		paths = append(paths, DefaultPath)
	}

	return &Params{
		Paths:    paths,
		Filename: DefaultFilename,
	}
}
