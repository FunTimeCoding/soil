package example

import (
	"example/inner"
	"example/pkg/assert"
)

func ExpectedLiteral() {
	assert.Any(nil, &inner.MyStruct{}, new(inner.MyStruct))
}
