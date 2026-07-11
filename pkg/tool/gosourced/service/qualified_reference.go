package service

import "github.com/funtimecoding/soil/pkg/source/resolve"

type qualifiedReference struct {
	reference resolve.Reference
	newName   string
}
