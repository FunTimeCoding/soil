package web

import (
	"github.com/funtimecoding/soil/pkg/errors/caller_left"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
)

func PanicOnWriteError(e error) {
	if e == nil {
		return
	}

	if f := connection.Classify(e); f != nil && f.Kind == errorsConstant.Dropped {
		panic(caller_left.New(e))
	}

	panic(e)
}
