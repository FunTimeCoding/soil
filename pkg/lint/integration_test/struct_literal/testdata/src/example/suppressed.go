package example

import "example/inner"

func Suppressed() *inner.MyStruct {
	return &inner.MyStruct{} // goanalyze:ignore struct_literal
}
