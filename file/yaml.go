package file

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func ParseYAML[T any](params *Params) (*T, error) {
	var cfg T
	filename := params.Filename + ".yaml"

	for _, path := range params.Paths {
		fullPath := filepath.Join(path, filename)

		_, err := os.Stat(fullPath)
		if err != nil {
			if os.IsNotExist(err) {
				continue // try next path
			}

			return nil, err // some other error occurred
		}

		yamlFile, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, err
		}

		if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
			return nil, err
		}

		return &cfg, nil
	}

	err := NewNotFound(filename, params.Paths)

	return nil, err
}
