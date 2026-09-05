package example

import "example/inner"

func UsesGenerated() *inner.GeneratedStruct {
	return &inner.GeneratedStruct{}
}
