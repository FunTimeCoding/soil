package output

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
)

type Results struct {
	workDirectory string
	Entries       []*concern.Concern
	Unchecked     []*unchecked.Unchecked
}
