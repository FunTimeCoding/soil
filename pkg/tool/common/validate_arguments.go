package common

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
)

func ValidateArguments(a *argument.Instance) {
	a.Required(constant.Host)
	a.Required(constant.Token)
	a.Required(constant.Owner)
	a.Required(constant.Repository)
}
