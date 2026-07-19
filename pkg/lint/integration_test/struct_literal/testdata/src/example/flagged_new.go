package example

import "example/inner"

func FlaggedNew() *inner.MyStruct {
	return new(inner.MyStruct)
}
