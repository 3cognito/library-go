package utils

import (
	"encoding/json"
	"errors"
)

func ConvertStruct(data interface{}, result interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return err
	}

	return nil
}

func ConvertDataToMap(data interface{}) map[string]interface{} {
	var f interface{}
	ConvertStruct(&data, &f)

	return f.(map[string]interface{})
}

func ConvertDataToMapWithError(data interface{}) (map[string]interface{}, error) {
	var f interface{}
	if err := ConvertStruct(&data, &f); err != nil {
		return nil, err
	}

	result, ok := f.(map[string]interface{})
	if !ok {
		return nil, errors.New("data is not a valid struct")
	}

	return result, nil
}
