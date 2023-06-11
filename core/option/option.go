package xzoption

import "reflect"

type Option[T any] func(value *T)

func SetValue[T any](st any, field string, value T) {
	if v := reflect.ValueOf(st).Elem().FieldByName(field); v.CanSet() {
		v.Set(reflect.ValueOf(value))
	}
}

func SetPrivateValue[T any](st any, methodName string, value T) {
	if fn := reflect.ValueOf(st).MethodByName(methodName); fn.IsValid() {
		fn.Call([]reflect.Value{reflect.ValueOf(value)})
	}
}
