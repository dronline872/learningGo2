package main

import (
	"fmt"
	"reflect"
)

type NewStruct struct {
	Id       int
	Name     string
	LastName string
}

func main() {
	NewMap := map[string]interface{}{
		"Id":       1,
		"Name":     "John",
		"LastName": "Doe",
	}

	NewStruct := NewStruct{}

	generateStruct(&NewStruct, NewMap)

	fmt.Println(NewStruct)
}

func generateStruct(in interface{}, values map[string]interface{}) error {
	if in == nil {
		return nil
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		valueStruct := val.FieldByName(typeField.Name)
		if valueMap, ok := values[typeField.Name].(int); ok {
			valueStruct.SetInt(int64(valueMap))
		} else if valueMap, ok := values[typeField.Name].(string); ok {
			valueStruct.SetString(valueMap)
		} else {
			continue
		}
	}

	return nil
}
