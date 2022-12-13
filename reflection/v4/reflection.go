package main

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	value := reflect.ValueOf(x)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Interface:
			fmt.Println(field)
			walk(field.Interface(), fn)
		}
	}
}

func walk2(x interface{}, fn func()) {
	value := reflect.ValueOf(x)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fmt.Println(field.Kind())
	}
}
