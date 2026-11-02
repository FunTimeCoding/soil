package relocation

import (
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"go/types"
)

func NewFieldFlip(
	object types.Object,
	newName string,
	references []resolve_reference.Reference,
) *FieldFlip {
	return &FieldFlip{Object: object, NewName: newName, References: references}
}
