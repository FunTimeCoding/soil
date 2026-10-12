package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/integers"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (i *Instance) Address() string {
	return join.Empty(
		i.GetString(constant.BindAddress),
		stringsConstant.Colon,
		integers.ToString(i.GetInteger(constant.Port)),
	)
}
