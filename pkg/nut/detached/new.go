package detached

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/nut/constant"
)

func New(
	battery string,
	code string,
) *DetachedError {
	return &DetachedError{
		Message: fmt.Sprintf(constant.DetachedFormat, battery, code),
	}
}
