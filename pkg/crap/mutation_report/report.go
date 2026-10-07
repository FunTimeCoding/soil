package mutation_report

import "github.com/funtimecoding/soil/pkg/crap/mutation"

type Report struct {
	Module string           `json:"go_module"`
	Files  []*mutation.File `json:"files"`
}
