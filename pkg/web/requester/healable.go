package requester

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
)

func healable(e error) bool {
	f := connection.Classify(e)

	if f == nil {
		return false
	}

	return f.Kind == errorsConstant.Timeout || f.Kind == errorsConstant.Dropped ||
		f.Reason == errorsConstant.Refused
}
