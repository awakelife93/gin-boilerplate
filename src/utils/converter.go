package utils

import (
	"fmt"
	"strconv"
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
