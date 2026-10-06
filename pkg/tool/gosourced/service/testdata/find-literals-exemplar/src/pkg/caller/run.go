package caller

import (
	"example/pkg/target"
	"fmt"
)

func multiline() *target.Shape {
	s := &target.Shape{
		Draw: 1,
		Label: fmt.Sprintf(
			"%s-%d",
			"alfa",
			2,
		),
	}

	return s
}
