package rename_plan

import (
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
	"go/types"
)

func New(
	rename *symbol_rename.Rename,
	declaration types.Object,
	references []resolve_reference.Reference,
) *Plan {
	return &Plan{
		Rename:      rename,
		Declaration: declaration,
		References:  references,
	}
}
