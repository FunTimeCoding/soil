package relocation

import "github.com/funtimecoding/soil/pkg/source/types/resolve_reference"

type QualifiedReference struct {
	Reference resolve_reference.Reference
	NewName   string
}
