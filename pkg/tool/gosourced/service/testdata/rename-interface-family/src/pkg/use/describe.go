package use

import (
	"example/pkg/box"
	"fmt"
)

func Describe() string {
	return fmt.Sprint(&box.Box{})
}
