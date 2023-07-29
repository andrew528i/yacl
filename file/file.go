package file

import (
	"os"
)

const DefaultFilename = "config"

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

	return &Params{
		Paths:    paths,
		Filename: DefaultFilename,
	}
}
