package reflection

import (
	"reflect"
	"unsafe"
)

// 类型转换
func Cast(slice interface{}, nArrayType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice || nArrayType.Kind() != reflect.Slice {
		return nil
	}
	newSlice := reflect.New(nArrayType)
	hd := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hd.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(nArrayType.Elem().Size())
	hd.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(nArrayType.Elem().Size())
	hd.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}
