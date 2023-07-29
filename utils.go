package yacl

import (
	"bytes"
	"encoding/gob"
)

func deepCopyStruct[T any](src *T) (*T, error) {
	// Create an encoder and decoder for Gob serialization.
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	// Encode the source struct.
	if err := enc.Encode(src); err != nil {
		return nil, err
	}

	// Decode the encoded struct into a new struct.
	var dst T
	if err := dec.Decode(&dst); err != nil {
		return nil, err
	}

	return &dst, nil
}
