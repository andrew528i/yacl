package flags

import (
	"testing"
)

func TestFromFlags_UnsupportedTypes(t *testing.T) {
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
				_ = FromFlags[FirstStruct](DefaultFlagParams())
			},
		},
		{
			name: "float32-slice",
			exec: func() {
				_ = FromFlags[SecondStruct](DefaultFlagParams())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("FromFlags did not panic as expected")
				}
			}()

			tc.exec()
		})
	}
}
