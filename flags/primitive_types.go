package flags

import (
	"flag"
	"reflect"
	"unsafe"
)

func bindStringFlag(flagName string, value reflect.Value) {
	flag.StringVar(value.Addr().Interface().(*string), flagName, "", "")
}

func bindBoolFlag(flagName string, value reflect.Value) {
	flag.BoolVar(value.Addr().Interface().(*bool), flagName, false, "")
}

func bindUintFlag(flagName string, value reflect.Value) {
	flag.UintVar((*uint)(unsafe.Pointer(value.Addr().Pointer())), flagName, 0, "")
}

func bindIntFlag(flagName string, value reflect.Value) {
	flag.IntVar((*int)(unsafe.Pointer(value.Addr().Pointer())), flagName, 0, "")
}

func bindFloat64Flag(flagName string, value reflect.Value) {
	flag.Float64Var((*float64)(unsafe.Pointer(value.Addr().Pointer())), flagName, .0, "")
}
