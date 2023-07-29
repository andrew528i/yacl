package flags

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func bindFloat64SliceFlag(flagName string, value reflect.Value) {
	flag.Var(newFloat64SliceValue(value), flagName, "")
}

type float64SliceValue struct {
	value reflect.Value
}

func newFloat64SliceValue(v reflect.Value) *float64SliceValue {
	return &float64SliceValue{v}
}

func (s *float64SliceValue) String() string {
	return fmt.Sprintf("%v", s.value.Interface())
}

func (s *float64SliceValue) Set(value string) error {
	f, err := strconv.ParseFloat(value, s.value.Type().Elem().Bits())
	if err != nil {
		return err
	}
	s.value.Set(reflect.Append(s.value, reflect.ValueOf(f)))
	return nil
}
