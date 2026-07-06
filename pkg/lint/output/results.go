package output

import "github.com/funtimecoding/soil/pkg/lint/concern"

type Results struct {
	workDirectory string
	Entries       []*concern.Concern
}
