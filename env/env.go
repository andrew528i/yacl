package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/andrew528i/yacl/utils"
)

var DefaultPrefix = ""
var DefaultDelimiter = "_"

type Params struct {
	Prefix    string
	Delimiter string
}

func DefaultParams() *Params {
	return &Params{
		Prefix:    DefaultPrefix,
		Delimiter: DefaultDelimiter,
	}
}

func Parse[T any](params *Params) (*T, error) {
	var cfg T
	names := make([]string, 0)
	callback := func(fieldPath []string, value reflect.Value, tag *reflect.StructTag) error {
		var name string

		fieldPathCopy := make([]string, 0, len(fieldPath))

		for _, p := range fieldPath {
			fieldPathCopy = append(fieldPathCopy, utils.CamelCaseToSlice(p)...)
		}

		if tag != nil && tag.Get("yacl") != "" {
			// TODO: add description and default values parsing from tag
			name = tag.Get("yacl")
		} else {
			name = strings.Join(fieldPathCopy, params.Delimiter)
		}

		if params.Prefix != "" {
			name = fmt.Sprintf("%s%s%s", params.Prefix, params.Delimiter, name)
		}

		name = strings.ToUpper(name)
		envVal := os.Getenv(name)

		if envVal == "" {
			return nil
		}

		switch value.Kind() {
		case reflect.String:
			value.SetString(envVal)

		case reflect.Bool:
			val, err := strconv.ParseBool(envVal)
			if err != nil {
				return err
			}

			value.SetBool(val)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := strconv.ParseUint(envVal, 10, 64)
			if err != nil {
				return err
			}

			value.SetUint(val)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, err := strconv.ParseInt(envVal, 10, 64)
			if err != nil {
				return err
			}

			value.SetInt(val)

		case reflect.Float64:
			val, err := strconv.ParseFloat(envVal, 64)
			if err != nil {
				return err
			}

			value.SetFloat(val)

		case reflect.Slice:
			elemKind := value.Type().Elem().Kind()
			switch elemKind {
			case reflect.String:
				vals := strings.Split(envVal, ",")
				slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))

				for i, v := range vals {
					slice.Index(i).SetString(v)
				}

				value.Set(slice)

			case reflect.Bool:
				vals := strings.Split(envVal, ",")
				slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))

				for i, v := range vals {
					val, err := strconv.ParseBool(v)
					if err != nil {
						return err
					}

					slice.Index(i).SetBool(val)
				}

				value.Set(slice)

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				vals := strings.Split(envVal, ",")
				slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))

				for i, v := range vals {
					val, err := strconv.ParseUint(v, 10, 64)
					if err != nil {
						return err
					}

					slice.Index(i).SetUint(val)
				}

				value.Set(slice)

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				vals := strings.Split(envVal, ",")
				slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))

				for i, v := range vals {
					val, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						return err
					}

					slice.Index(i).SetInt(val)
				}

				value.Set(slice)

			case reflect.Float64:
				vals := strings.Split(envVal, ",")
				slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))

				for i, v := range vals {
					val, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return err
					}

					slice.Index(i).SetFloat(val)
				}

				value.Set(slice)

			default:
				panic(fmt.Sprintf("slice type not supported: `%s`", value.Type().Name()))
			}

		default:
			panic(fmt.Sprintf("type not supported: `%s`", value.Type().Name()))
		}

		names = append(names, name)

		return nil
	}

	if err := utils.WalkStruct(&cfg, callback); err != nil {
		return nil, err
	}

	return &cfg, nil
}
