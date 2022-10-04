package utils

import (
	"fmt"
	"strconv"

	"github.com/awakelife93/gin-boilerplate/src/config"
)

func GetStringValue(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

func GetStringToInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}

	_value, error := strconv.Atoi(value)

	if error != nil {
		fmt.Println("GetStringToInt Error", error.Error())
		_value = defaultValue
	}

	return _value
}

func GetRequestId(item interface{}) int {
	_item := item.(map[string]interface{})

	if len(_item) < 1 {
		fmt.Println("GetRequestId Error ========> item is empty.")
	}

	requestId := _item[config.RequestIdFieldName()].(string)
	return GetStringToInt(requestId, 0)
}
