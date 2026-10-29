package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/constant"
)

func directoryDown() error {
	return connection.New(
		constant.Unreachable,
		"directory.test",
		"",
		constant.Refused,
	)
}
