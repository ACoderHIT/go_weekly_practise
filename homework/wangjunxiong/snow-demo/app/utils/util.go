package utils

import (
	"reflect"
	"strconv"
)

func TypeTransfer(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
		break
	case int:
		return strconv.Itoa(value.(int))
		break
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', 0, 64)
		break
	}

	typeName := reflect.TypeOf(value).Name()
	if typeName == "string" {
		return value.(string)
	}
	if typeName == "float64" {
		return strconv.FormatFloat(value.(float64), 'E', -1, 64)
	}
	if typeName == "int64" {
		return strconv.FormatInt(value.(int64), 10)
	}
	return ""
}
