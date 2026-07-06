package argument

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/spf13/pflag"
)

func NewInstance(i *identity.Tool) *Instance {
	f := pflag.NewFlagSet(i.Name(), pflag.ContinueOnError)
	i.SetUsageOn(f)

	return &Instance{
		identity: i,
		flags:    f,
	}
}
