package relocation

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/types"
)

func NewFieldFlip(
	object types.Object,
	newName string,
	references []resolve.Reference,
) *FieldFlip {
	return &FieldFlip{Object: object, NewName: newName, References: references}
}
