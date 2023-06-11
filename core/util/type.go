package xzutil

import (
	"fmt"
	"reflect"
)

func IsPointer(result any) error {
	if reflect.ValueOf(result).Kind() != reflect.Pointer {
		return fmt.Errorf("解析数据应为指针类型")
	}
	return nil
}
