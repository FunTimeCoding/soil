package requester

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/constant"
)

func healable(e error) bool {
	f := connection.Classify(e)

	if f == nil {
		return false
	}

	return f.Kind == constant.Timeout || f.Kind == constant.Dropped ||
		f.Reason == constant.Refused
}
