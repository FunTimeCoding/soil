package example

import "example/inner"

func AllowedPointer() *inner.MyStruct {
	return &inner.MyStruct{Value: 1}
}

func AllowedPrimitive() int {
	return 42
}

func AllowedString() string {
	return "hello"
}

func AllowedError() error {
	return nil
}

func AllowedSlice() []*inner.MyStruct {
	return nil
}
