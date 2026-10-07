package virtual_file_system

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system/file"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system/pending_move"
)

type System struct {
	files    map[string]*file.File
	moves    []pending_move.Move
	written  map[string]bool
	deleted  map[string]bool
	maxFiles int
	maxBytes int64
}
