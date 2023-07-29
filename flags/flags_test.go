package flags

import (
	"testing"
)

func TestParse_UnsupportedTypes(t *testing.T) {
	type FirstStruct struct {
		Value float32
	}

	type SecondStruct struct {
		Values []float32
	}

	testCases := []struct {
		name string
		exec func()
	}{
		{
			name: "float32",
			exec: func() {
				_ = Parse[FirstStruct](DefaultFlagParams())
			},
		},
		{
			name: "float32-slice",
			exec: func() {
				_ = Parse[SecondStruct](DefaultFlagParams())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("Parse did not panic as expected")
				}
			}()

			tc.exec()
		})
	}
}
