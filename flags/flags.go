package flags

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/andrew528i/yacl/utils"
)

type FlagParams struct {
	Delimiter           string
	FieldPathFormatFunc func([]string)
}

func DefaultFlagParams() *FlagParams {
	delimiter := "-"

	return &FlagParams{
		Delimiter: delimiter,
		FieldPathFormatFunc: func(fieldPath []string) {
			for i := 0; i < len(fieldPath); i++ {
				fieldName := strings.Join(camelCaseToSlice(fieldPath[i]), delimiter)
				fieldPath[i] = strings.ToLower(fieldName)
			}
		},
	}
}

func FromFlags[T any](params *FlagParams) *T {
	var cfg T

	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

	utils.WalkStruct[T](&cfg, func(fieldPath []string, value reflect.Value, tag *reflect.StructTag) {
		var flagName string
		fieldPathCopy := make([]string, len(fieldPath))
		copy(fieldPathCopy, fieldPath)

		if tag != nil && tag.Get("yacl") != "" {
			// TODO: add description and default values parsing from tag
			flagName = tag.Get("yacl")
		} else {
			if params.FieldPathFormatFunc != nil {
				params.FieldPathFormatFunc(fieldPathCopy)
			}

			flagName = strings.Join(fieldPathCopy, params.Delimiter)
		}

		switch value.Kind() {
		case reflect.String:
			bindStringFlag(flagName, value)

		case reflect.Bool:
			bindBoolFlag(flagName, value)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			bindUintFlag(flagName, value)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			bindIntFlag(flagName, value)

		case reflect.Float64:
			bindFloat64Flag(flagName, value)

		case reflect.Slice:
			elemKind := value.Type().Elem().Kind()
			switch elemKind {
			case reflect.String:
				bindStringSliceFlag(flagName, value)

			case reflect.Bool:
				bindBoolSliceFlag(flagName, value)

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				bindUintSliceFlag(flagName, value)

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				bindIntSliceFlag(flagName, value)

			case reflect.Float64:
				bindFloat64SliceFlag(flagName, value)

			default:
				panic(fmt.Sprintf("slice type not supported: `%s`", value.Type().Name()))
			}

		default:
			panic(fmt.Sprintf("type not supported: `%s`", value.Type().Name()))
		}
	})

	flag.Parse()

	return &cfg
}
