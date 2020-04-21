package dyl

import (
	"reflect"
)

func Option(t interface{}, o interface{}) reflect.Value {
	xs := reflect.ValueOf(o)
	tv := reflect.ValueOf(t)
	for i := 0; i < xs.Len(); i++ {
		x := xs.Index(i)
		if x.Kind() == reflect.Interface {
			x = x.Elem()
		}
		if x.Type() == tv.Type() {
			return x
		}
	}
	return tv
}

func IfsOption(t interface{}, o []interface{}) interface{} {
	return Option(t, o).Interface()
}

func MultiOption(o []interface{}, t ...interface{}) (reflect.Value, int) {
	for _, x := range o {
		for i, tv := range t {
			v := reflect.ValueOf(x)
			if v.Type() == reflect.TypeOf(tv) {
				return v, i
			}
		}
	}
	return reflect.ValueOf(t[0]), 0
}

func StrMultiOption(o []interface{}, t ...interface{}) (string, int) {
	v, i := MultiOption(o, t...)
	return v.String(), i
}
