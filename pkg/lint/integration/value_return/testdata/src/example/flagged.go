package example

import "example/inner"

func Flagged() inner.MyStruct {
	return inner.MyStruct{Value: 1}
}

func FlaggedWithError() (inner.MyStruct, error) {
	return inner.MyStruct{}, nil
}
