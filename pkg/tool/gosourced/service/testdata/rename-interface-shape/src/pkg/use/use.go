package use

import (
	"example/pkg/door"
	"example/pkg/shape"
)

func Use() error {
	var c shape.Closer = &door.Door{}

	return c.Close()
}
