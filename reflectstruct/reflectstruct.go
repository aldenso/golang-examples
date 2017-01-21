package main

import (
	"fmt"

	"reflect"
)

//Request some example flat struct.
type Request struct {
	Name  string  `json:"name"`
	ID    int     `json:"id"`
	Email string  `json:"email"`
	Score float64 `json:"score"`
}

//GetInfo method to get tags, fields and types from flat struct.
func (r *Request) GetInfo() ([]string, []string, []string) {
	var tags, fields, types []string
	val := reflect.ValueOf(r).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		valuestring := fmt.Sprintf("%v", valueField.Interface())
		fields = append(fields, valuestring)
		typefield := val.Type().Field(i)
		types = append(types, typefield.Name)
		tag := typefield.Tag
		tags = append(tags, tag.Get("json"))

	}
	return tags, fields, types
}

func main() {
	request := &Request{Name: "aldenso", ID: 1, Email: "aldenso@gmail.com", Score: 50.50}
	tags, fields, types := request.GetInfo()
	fmt.Printf("slice of tags: %v\nslice of fields: %v\nslice of types: %v\n", tags, fields, types)
	fmt.Printf("original request: %v\n", *request)
}
