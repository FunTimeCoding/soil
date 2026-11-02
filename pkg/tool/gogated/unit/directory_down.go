package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
)

func directoryDown() error {
	return connection.New(
		errorsConstant.Unreachable,
		"directory.test",
		"",
		errorsConstant.Refused,
	)
}
