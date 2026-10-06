package server

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/stamp"
)

func New(i *identity.Tool) *Builder {
	return &Builder{
		name:         i.Name(),
		version:      stamp.New().Tag(),
		instructions: i.Instructions(),
	}
}
