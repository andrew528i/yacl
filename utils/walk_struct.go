package utils

import "reflect"

type WalkStructCallback func(fieldPath []string, field reflect.Value, tag *reflect.StructTag) error

func WalkStruct[T any](s *T, callback WalkStructCallback) error {
	value := reflect.ValueOf(s).Elem()

	return walkStruct(value, []string{}, callback, nil)
}

func walkStruct(value reflect.Value, fields []string, callback WalkStructCallback, tag *reflect.StructTag) error {
	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldType := value.Type().Field(i)
			fieldName := fieldType.Name
			fieldTag := fieldType.Tag

			if field.Kind() == reflect.Struct {
				if err := walkStruct(field, append(fields, fieldName), callback, &fieldTag); err != nil {
					return err
				}
			} else {
				if err := callback(append(fields, fieldName), field, &fieldTag); err != nil {
					return err
				}
			}
		}

	default:
		if err := callback(fields, value, tag); err != nil {
			return err
		}
	}

	return nil
}
