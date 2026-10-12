package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/integers"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (i *Instance) MetricAddress() string {
	return join.Empty(
		stringsConstant.Colon,
		integers.ToString(i.GetInteger(constant.MetricPort)),
	)
}
