package server

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
)

func isClientError(e error) bool {
	return not_found.Is(e) || ambiguous.Is(e) || not_selected.Is(e)
}
