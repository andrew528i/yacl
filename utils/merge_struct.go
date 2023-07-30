package utils

import "reflect"

func MergeStruct(dst, src interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < dstValue.NumField(); i++ {
		dstField := dstValue.Field(i)
		srcField := srcValue.Field(i)

		if srcField.Kind() != dstField.Kind() {
			continue
		}

		switch srcField.Kind() {
		case reflect.Slice:
			if !srcField.IsNil() {
				dstField.Set(srcField)
			}

		case reflect.Struct:
			MergeStruct(dstField.Addr().Interface(), srcField.Addr().Interface())

		default:
			if srcField.Interface() != reflect.Zero(srcField.Type()).Interface() {
				dstField.Set(srcField)
			}
		}
	}
}
