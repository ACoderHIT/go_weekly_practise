package input_util

import (
	"fmt"
	"github.com/gogap/errors"
	"reflect"
	"strconv"
	"context"
)

var (
	ErrFieldNotExist = fmt.Errorf("field is not exist ")
	errValueIsNullTpl = errors.TN("input_util", 4001, "params:{{.keyName}} value is not valid")
)

// 获取int值
func GetIntValueByNameFromMap(ctx context.Context, inputData map[string]interface{}, keyName string, defaultVal int) (int, error) {
	if dataTemp, ok := inputData[keyName]; ok {
		dataTempV := reflect.ValueOf(dataTemp)
		if !dataTempV.IsValid() { // null
			// return 0, errValueIsNullTpl.New(errors.Params{"keyName": keyName})
			return defaultVal, nil
		}

		switch dataTempV.Type().Kind() {
		case reflect.Float64, reflect.Float32:
			idFloat64Temp, ok := dataTemp.(float64)
			if !ok {
				return 0, fmt.Errorf("params:%s convert to float64 failure", keyName)
			}
			return int(idFloat64Temp), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			idIntTemp, ok := dataTemp.(int)
			if !ok {
				return 0, fmt.Errorf("params:%s convert to int failure", keyName)
			}
			return idIntTemp, nil

		case reflect.Int64, reflect.Uint64:
			idIntTemp, ok := dataTemp.(int64)
			if !ok {
				return 0, fmt.Errorf("params:%s convert to int64 failure", keyName)
			}
			return int(idIntTemp), nil
		case reflect.String:
			idIntTemp, err := strconv.Atoi(dataTemp.(string))
			if err != nil {
				return 0, err
			}
			return idIntTemp, nil
		default:
			return 0, fmt.Errorf("params:%s's type is not allow.(string/float/int)", keyName)
		}

	}

	return defaultVal, nil
}

// 获取 int 列表
func GetIntValuesByNameFromMap(ctx context.Context, inputData map[string]interface{}, keyName string, defaultValues []int) ([]int, error) {

	var idsReturn = make([]int, 0)

	if dataTemp, ok := inputData[keyName]; ok {
		dataTempV := reflect.ValueOf(dataTemp)
		if !dataTempV.IsValid() { // null
			// return nil, errValueIsNullTpl.New(errors.Params{"keyName": keyName})
			return defaultValues, nil
		}

		switch dataTempV.Type().Kind() {
		case reflect.Array, reflect.Slice:
			isExistInvalid := false
			for _, dataItemTemp := range dataTemp.([]interface{}) {
				dataItemTempV := reflect.ValueOf(dataItemTemp)
				if !dataItemTempV.IsValid() { // null
					if !isExistInvalid {
						isExistInvalid = true
					}
					continue
				}

				switch dataItemTempV.Type().Kind() {
				case reflect.Float64, reflect.Float32:
					idFloat64Temp, ok := dataItemTemp.(float64)
					if !ok {
						return nil, fmt.Errorf("params:%s convert to float64 failure", keyName)
					}
					idIntTemp := int(idFloat64Temp)
					idsReturn = append(idsReturn, idIntTemp)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
					idIntTemp, ok := dataItemTemp.(int)
					if !ok {
						return nil, fmt.Errorf("params:%s convert to int failure", keyName)
					}
					idsReturn = append(idsReturn, idIntTemp)
				case reflect.Int64, reflect.Uint64:
					idIntTemp, ok := dataTemp.(int64)
					if !ok {
						return nil, fmt.Errorf("params:%s convert to int64 failure", keyName)
					}
					idsReturn = append(idsReturn, int(idIntTemp))
				case reflect.String:
					idIntTemp, err := strconv.Atoi(dataItemTemp.(string))
					if err != nil {
						return nil, err
					}
					idsReturn = append(idsReturn, idIntTemp)
				}
			}
			if len(idsReturn) <= 0 && isExistInvalid {
				idsReturn = append(idsReturn, 0)
			}
		default:
			return nil, fmt.Errorf("params:%s is not an array", keyName)
		}

		return idsReturn, nil

	}
	return defaultValues, nil
}

// 获取 string 列表
func GetStringValuesByNameFromMap(ctx context.Context, inputData map[string]interface{}, keyName string, defaultValues []string) ([]string, error) {

	var strListReturn = make([]string, 0)

	if dataTemp, ok := inputData[keyName]; ok {
		dataTempV := reflect.ValueOf(dataTemp)
		if !dataTempV.IsValid() { // null
			// return nil, errValueIsNullTpl.New(errors.Params{"keyName": keyName})
			return defaultValues,nil
		}

		switch dataTempV.Type().Kind() {
		case reflect.Array, reflect.Slice:
			for _, dataItemTemp := range dataTemp.([]interface{}) {
				dataItemTempV := reflect.ValueOf(dataItemTemp)
				if !dataItemTempV.IsValid() { // null
					continue
				}

				switch dataItemTempV.Type().Kind() {
				case reflect.String:
					strListReturn = append(strListReturn, dataItemTemp.(string))
				}
			}
		default:
			return nil, fmt.Errorf("params:%s is not an array", keyName)
		}

		return strListReturn, nil

	}
	return defaultValues, nil
}
