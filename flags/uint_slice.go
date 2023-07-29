package flags

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func bindUintSliceFlag(flagName string, value reflect.Value) {
	flag.Var(newUintSliceValue(value), flagName, "")
}

type uintSliceValue struct {
	value reflect.Value
}

func newUintSliceValue(v reflect.Value) *uintSliceValue {
	return &uintSliceValue{v}
}

func (s *uintSliceValue) String() string {
	return fmt.Sprintf("%v", s.value.Interface())
}

func (s *uintSliceValue) Set(value string) error {
	u, err := strconv.ParseUint(value, 0, s.value.Type().Elem().Bits())
	if err != nil {
		return err
	}
	s.value.Set(reflect.Append(s.value, reflect.ValueOf(u).Convert(s.value.Type().Elem())))
	return nil
}
