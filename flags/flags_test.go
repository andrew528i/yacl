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
				_, _ = Parse[FirstStruct](DefaultParams())
			},
		},
		{
			name: "float32-slice",
			exec: func() {
				_, _ = Parse[SecondStruct](DefaultParams())
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
