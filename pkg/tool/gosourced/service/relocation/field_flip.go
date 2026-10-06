package relocation

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/types"
)

type FieldFlip struct {
	Object     types.Object
	NewName    string
	References []resolve.Reference
}
