package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/integers"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (i *Instance) MetricAddress() string {
	return join.Empty(
		strings.Colon,
		integers.ToString(i.GetInteger(constant.MetricPort)),
	)
}
