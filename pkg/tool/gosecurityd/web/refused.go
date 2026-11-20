package web

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func refused(e error) bool {
	return validation.Is(e) || conflict.Is(e) || not_found.Is(e)
}
