package flags

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func bindBoolSliceFlag(flagName string, value reflect.Value) {
	flag.Var(newBoolSliceValue(value), flagName, "")
}

type boolSliceValue struct {
	value reflect.Value
}

func newBoolSliceValue(v reflect.Value) *boolSliceValue {
	return &boolSliceValue{v}
}

func (s *boolSliceValue) String() string {
	return fmt.Sprintf("%v", s.value.Interface())
}

func (s *boolSliceValue) Set(value string) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	s.value.Set(reflect.Append(s.value, reflect.ValueOf(b)))
	return nil
}
