package check_tester

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	gosourcedConstant "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"strconv"
)

func (o *Tester) ServeSource(p ...base.Option) {
	o.t.Helper()
	s := base.New(o.t, p...).Served()
	o.t.Setenv(gosourcedConstant.HostEnvironment, "localhost")
	o.t.Setenv(gosourcedConstant.PortEnvironment, strconv.Itoa(s.Port))
	o.t.Setenv(gosourcedConstant.InsecureEnvironment, "true")
	o.t.Setenv(
		gosourcedConstant.TokenEnvironment,
		generativeConstant.ModelContextTestToken,
	)
}
