package retargeted_document

import "github.com/funtimecoding/soil/pkg/lint/types/retarget_change"

type Document struct {
	Path     string
	Original string
	Content  string
	Changes  []*retarget_change.Change
}
