package relocation

import "github.com/funtimecoding/soil/pkg/source/resolve"

type QualifiedReference struct {
	Reference resolve.Reference
	NewName   string
}
