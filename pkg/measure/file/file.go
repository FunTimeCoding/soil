package file

import "github.com/funtimecoding/soil/pkg/measure/count"

type File struct {
	Path     string       `json:"path"`
	Language string       `json:"language"`
	Count    *count.Count `json:"count"`
}
