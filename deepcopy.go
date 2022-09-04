package main

import (
	"reflect"
	"unsafe"
)

func Copy(any interface{}) interface{} {
	val := reflect.ValueOf(any)
	res := copyValue(val)
	return res.Interface()
}

func copyValue(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Ptr:
		value := reflect.New(v.Elem().Type())
		memmove(value.UnsafePointer(), v.UnsafePointer(), value.Elem().Type().Size())
		if value.Elem().Kind() == reflect.Struct {
			res := copyValueFields(value.Elem())
			return res.Addr()
		}
		return value
	case reflect.Map:
		newMap := reflect.MakeMapWithSize(v.Type(), v.Len())
		for _, k := range v.MapKeys() {
			key := copyValue(k)
			value := copyValue(v.MapIndex(k))
			newMap.SetMapIndex(key, value)
		}
		return newMap
	case reflect.Struct:
		return copyValueFields(v)
	case reflect.Slice:
		slice := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		return copyArraySlice(slice, v)
	case reflect.Array:
		array := reflect.New(v.Type()).Elem()
		return copyArraySlice(array, v)
	default:
	}

	return v
}

func copyArraySlice(target, src reflect.Value) reflect.Value {
	reflect.Copy(target, src)
	for i := 0; i < src.Len(); i++ {
		val := target.Index(i)
		val.Set(copyValue(val))
	}
	return target
}

func copyValueFields(value reflect.Value) reflect.Value {
	res := reflect.New(value.Type()).Elem()
	for i := 0; i < value.NumField(); i++ {
		v := value.Field(i)
		val := copyValue(v)
		res.Field(i).Set(val)
	}
	return res
}

//go:linkname memmove runtime.memmove
func memmove(d unsafe.Pointer, s unsafe.Pointer, size uintptr)
