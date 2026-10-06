package use

import (
	"example/pkg/box"
	"example/pkg/shape"
	"fmt"
)

func Use() string {
	b := &box.Box{}
	var s shape.Sizer = b

	return fmt.Sprint(b, s.Size(), b.Weight())
}
