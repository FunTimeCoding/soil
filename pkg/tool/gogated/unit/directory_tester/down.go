package directory_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
)

func Down() error {
	return connection.New(
		errorsConstant.Unreachable,
		"directory.test",
		"",
		errorsConstant.Refused,
	)
}
