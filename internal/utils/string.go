package utils

import (
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase ...
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ToUpperSnakeCase ...
func ToUpperSnakeCase(str string) string {
	return strings.ToUpper(ToSnakeCase(str))
}

// NilToEmpty ...
func NilToEmpty(s *string) string {
	if s == nil {
		temp := ""
		return temp
	}

	return *s
}

// GetField ...
func GetField(i interface{}, fieldName string, newVal interface{}) {
	v := reflect.ValueOf(i)
	// Get the first element of the slice.
	e := v.Index(0)
	// Get the field of the slice element that we want to set.
	f := e.FieldByName(fieldName)
	// Set the value!
	f.Set(reflect.ValueOf(newVal))
}
