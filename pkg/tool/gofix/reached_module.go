package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"go/token"
)

type reachedModule struct {
	directory string
	fileSet   *token.FileSet
	edits     []edit
	concerns  []*concern.Concern
}
