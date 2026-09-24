package result

import "github.com/funtimecoding/soil/pkg/measure/file"

type Result struct {
	Files    []*file.File `json:"files"`
	Unplaced []string     `json:"unplaced"`
	Skipped  []string     `json:"skipped"`
}
