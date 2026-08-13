package argument

import "github.com/spf13/pflag"

func NewSimple(name string) *Instance {
	return &Instance{flags: pflag.NewFlagSet(name, pflag.ContinueOnError)}
}
