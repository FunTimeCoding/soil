package use

import (
	"example/pkg/box"
	"example/pkg/mock"
	"example/pkg/shape"
)

func Use() int {
	b := &box.Box{}
	var s shape.Sizer = b
	var m shape.Measured = &mock.Mock{}

	return s.Size() + m.Size() + b.Size()
}
