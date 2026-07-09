package argument

import (
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
)

func (i *Instance) Address() string {
	return join.Empty(
		i.GetString(BindAddress),
		separator.Colon,
		integers.ToString(i.GetInteger(Port)),
	)
}
