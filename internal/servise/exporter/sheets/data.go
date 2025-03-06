package sheets

import (
	"fmt"
	"reflect"
	"strings"
)

type data []sheetsRow

type sheetsRow []interface{}

func parseData(d [][]interface{}) data {
	out := make(data, len(d))
	for i, r := range d {
		out[i] = r
	}
	return out
}

func (d data) Len() int {
	return len(d)
}

func (r sheetsRow) Unmarshal(dst interface{}) error {
	reflectV := reflect.ValueOf(dst)
	if reflectV.Kind() != reflect.Ptr || reflectV.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dst must be a pointer to a struct")
	}
	reflectV = reflectV.Elem()
	for i, rawCeil := range r {
		field := reflectV.Field(i)
		switch field.Kind() {
		case reflect.String:
			strCeil, ok := rawCeil.(string)
			if !ok {
				return fmt.Errorf("field %d must be a string", i)
			}
			field.SetString(strCeil)
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64, reflect.Uint,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			floatCeil, ok := rawCeil.(float64)
			if !ok {
				return fmt.Errorf("cannot convert %v to int", rawCeil)
			}
			field.SetInt(int64(floatCeil))
		case reflect.Float32, reflect.Float64:
			floatCeil, ok := rawCeil.(float64)
			if !ok {
				return fmt.Errorf("cannot convert %v to int", rawCeil)
			}
			field.SetFloat(floatCeil)
		case reflect.Bool:
			boolCeil, ok := rawCeil.(bool)
			if !ok {
				return fmt.Errorf("cannot convert %v to int", rawCeil)
			}
			field.SetBool(boolCeil)
		case reflect.Slice:
			elemKind := field.Type().Elem().Kind()
			switch elemKind {
			case reflect.String:
				switch v := rawCeil.(type) {
				case string:
					field.Set(
						reflect.ValueOf(
							strings.Split(strings.TrimSpace(v), ", "),
						),
					)
				case []interface{}:
					strSlice := make([]string, len(v))
					for i, val := range v {
						strVal, ok := val.(string)
						if !ok {
							return fmt.Errorf("cannot convert %v to slice", val)
						}
						strSlice[i] = strVal
					}
					field.Set(reflect.ValueOf(strSlice))
				default:
					return fmt.Errorf("cannot convert %v to slice", rawCeil)
				}
			case reflect.Uint8:
				switch v := rawCeil.(type) {
				case string:
					field.Set(reflect.ValueOf([]byte(v)))
				case []interface{}:
					byteSlice := make([]byte, len(v))
					for i, val := range v {
						num, ok := val.(float64)
						if !ok || num < 0 || num > 255 {
							return fmt.Errorf("cannot convert %v to slice", val)
						}
						byteSlice[i] = byte(num)
					}
					field.Set(reflect.ValueOf(byteSlice))
				default:
					return fmt.Errorf("cannot convert %v to slice", rawCeil)
				}

			default:
				return fmt.Errorf("cannot convert %v to slice", rawCeil)
			}
		default:
			return fmt.Errorf("cannot unmarshal field %s", field.Kind())
		}
	}
	return nil
}
