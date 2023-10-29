package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertStringToType(value string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		return intValue, err
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		return boolValue, err
	case reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		return floatValue, err
	case reflect.String:
		return value, nil
	default:
		return nil, fmt.Errorf("Unsupported kind: %v", kind)
	}
}
