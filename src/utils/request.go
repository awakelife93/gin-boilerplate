package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/awakelife93/gin-boilerplate/src/config"
)

func GenerateBodyItem(body io.ReadCloser) interface{} {
	var result interface{}

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	error := decoder.Decode(&result)

	if error != nil {
		fmt.Println("Request DecodeBody Error ========>", error.Error())
		return nil
	}

	return result
}

func GenerateQueryItem(query url.Values, params string) map[string]interface{} {
	var result = make(map[string]interface{})

	if params != "" {
		result[config.RequestIdFieldName()] = params
	}

	for key := range query {
		result[key] = query.Get(key)
	}

	return result
}
