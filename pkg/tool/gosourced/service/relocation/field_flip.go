package relocation

import (
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"go/types"
)

type FieldFlip struct {
	Object     types.Object
	NewName    string
	References []resolve_reference.Reference
}
