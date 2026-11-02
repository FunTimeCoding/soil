package service

import (
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
	"go/types"
)

type plannedRename struct {
	Rename      *symbol_rename.Rename
	Declaration types.Object
	References  []resolve_reference.Reference
}
