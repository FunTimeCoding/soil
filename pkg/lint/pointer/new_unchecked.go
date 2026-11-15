package pointer

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewUnchecked(reason constant.Reason) *Resolution {
	return &Resolution{Reason: reason}
}
