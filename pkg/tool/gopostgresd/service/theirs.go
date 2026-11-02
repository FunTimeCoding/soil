package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/caller_left"
)

func theirs(
	x context.Context,
	e error,
) error {
	if x.Err() != nil {
		return caller_left.New(e)
	}

	if timeout.Is(e) {
		return validation.New("query ran too long: %s", e)
	}

	return e
}
