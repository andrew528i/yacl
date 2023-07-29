package flags

import (
	"flag"
	"fmt"
	"reflect"
)

func bindStringSliceFlag(flagName string, value reflect.Value) {
	flag.Var(newStringSliceValue(value), flagName, "")
}

type stringSliceValue struct {
	value reflect.Value
}

func newStringSliceValue(v reflect.Value) *stringSliceValue {
	return &stringSliceValue{v}
}

func (s *stringSliceValue) String() string {
	return fmt.Sprintf("%v", s.value.Interface())
}

func (s *stringSliceValue) Set(value string) error {
	s.value.Set(reflect.Append(s.value, reflect.ValueOf(value)))

	return nil
}
