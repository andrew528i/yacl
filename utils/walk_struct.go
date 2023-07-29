package utils

import "reflect"

type WalkStructCallback func(fieldPath []string, field reflect.Value, tag *reflect.StructTag)

func WalkStruct[T any](s *T, callback WalkStructCallback) {
	value := reflect.ValueOf(s).Elem()

	walkStruct(value, []string{}, callback, nil)
}

func walkStruct(value reflect.Value, fields []string, callback WalkStructCallback, tag *reflect.StructTag) {
	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldType := value.Type().Field(i)
			fieldName := fieldType.Name
			fieldTag := fieldType.Tag

			if field.Kind() == reflect.Struct {
				walkStruct(field, append(fields, fieldName), callback, &fieldTag)
			} else {
				callback(append(fields, fieldName), field, &fieldTag)
			}
		}

	default:
		callback(fields, value, tag)
	}
}
