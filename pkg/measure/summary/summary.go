package summary

import "github.com/funtimecoding/soil/pkg/measure/count"

type Summary struct {
	Language string       `json:"language"`
	Files    int          `json:"files"`
	Count    *count.Count `json:"count"`
}
