package yacl

import (
	"github.com/andrew528i/yacl/env"
	"github.com/andrew528i/yacl/file"
	"github.com/andrew528i/yacl/flags"
)

func SetEnvPrefix(prefix string) {
	env.DefaultPrefix = prefix
}

func SetEnvDelimiter(delimiter string) {
	env.DefaultDelimiter = delimiter
}

func SetFilePath(path string) {
	file.DefaultPath = path
}

func SetFilename(filename string) {
	file.DefaultFilename = filename
}

func SetFlagDelimiter(delimiter string) {
	flags.DefaultDelimiter = delimiter
}
