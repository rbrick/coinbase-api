package coinbase

import (
	"reflect"
)

func FieldIndexByTag(data interface{}, tag string) (int, bool) {
	typeOf := reflect.TypeOf(data)

	if typeOf.Kind() == reflect.Struct || typeOf.Kind() == reflect.Ptr {
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
			if typeOf.Kind() != reflect.Struct {
				return -1, false
			}
		}

		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)

			if _, ok := field.Tag.Lookup(tag); ok {
				return i, ok
			}
		}

	}

	return -1, false

}

func Data(data interface{}, index int) reflect.Value {
	valueOf := reflect.ValueOf(data)

	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	f := valueOf.Field(index)

	return f
}

func SetField(i, val interface{}, index int) {
	valueOf := reflect.ValueOf(i)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	valueOf.Field(index).Set(reflect.ValueOf(val))
}
