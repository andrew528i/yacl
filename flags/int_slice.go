package flags

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func bindIntSliceFlag(flagName string, value reflect.Value) {
	flag.Var(newIntSliceValue(value), flagName, "")
}

type intSliceValue struct {
	value reflect.Value
}

func newIntSliceValue(v reflect.Value) *intSliceValue {
	return &intSliceValue{v}
}

func (s *intSliceValue) String() string {
	return fmt.Sprintf("%v", s.value.Interface())
}

func (s *intSliceValue) Set(value string) error {
	i, err := strconv.ParseInt(value, 0, s.value.Type().Elem().Bits())
	if err != nil {
		return err
	}
	s.value.Set(reflect.Append(s.value, reflect.ValueOf(i).Convert(s.value.Type().Elem())))
	return nil
}
