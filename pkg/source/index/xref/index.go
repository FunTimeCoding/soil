package xref

import "github.com/funtimecoding/soil/pkg/source/index/record"

type Index struct {
	units       map[string]*record.References
	directories map[string]string
}
