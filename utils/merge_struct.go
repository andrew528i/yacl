package utils

import "reflect"

func MergeStruct(dst, src interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < dstValue.NumField(); i++ {
		dstField := dstValue.Field(i)
		srcField := srcValue.Field(i)

		if srcField.Interface() != reflect.Zero(srcField.Type()).Interface() {
			if dstField.Kind() == reflect.Struct && srcField.Kind() == reflect.Struct {
				MergeStruct(dstField.Addr().Interface(), srcField.Addr().Interface())
			} else {
				dstField.Set(srcField)
			}
		}
	}
}
