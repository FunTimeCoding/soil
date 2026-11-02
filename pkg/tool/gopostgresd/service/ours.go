package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/caller_left"
)

func ours(
	x context.Context,
	e error,
) error {
	if x.Err() != nil {
		return caller_left.New(e)
	}

	if unavailable.Is(e) || classify.Reportable(e) {
		return e
	}

	return unexpected.Format("catalogue query refused: %s", e)
}
