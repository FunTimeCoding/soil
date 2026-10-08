package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"go/token"
)

type ReachedModule struct {
	directory string
	fileSet   *token.FileSet
	edits     []Edit
	concerns  []*concern.Concern
}
