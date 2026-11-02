package common

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
)

func ValidateArguments(a *argument.Instance) {
	a.Required(argumentConstant.Host)
	a.Required(argumentConstant.Token)
	a.Required(argumentConstant.Owner)
	a.Required(argumentConstant.Repository)
}
