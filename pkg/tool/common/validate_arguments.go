package common

import "github.com/funtimecoding/soil/pkg/argument"

func ValidateArguments(a *argument.Instance) {
	a.Required(argument.Host)
	a.Required(argument.Token)
	a.Required(argument.Owner)
	a.Required(argument.Repository)
}
