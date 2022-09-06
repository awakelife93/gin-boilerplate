package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
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

func GenerateQueryItem(query url.Values) map[string]interface{} {
	var result = make(map[string]interface{})

	for key := range query {
		result[key] = query.Get(key)
	}

	return result
}
