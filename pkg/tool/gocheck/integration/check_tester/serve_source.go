package check_tester

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	gosourcedConstant "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"strconv"
	"testing"
)

func ServeSource(
	t *testing.T,
	o ...base.Option,
) {
	t.Helper()
	s := base.New(t, o...).Served()
	t.Setenv(gosourcedConstant.HostEnvironment, "localhost")
	t.Setenv(gosourcedConstant.PortEnvironment, strconv.Itoa(s.Port))
	t.Setenv(gosourcedConstant.InsecureEnvironment, "true")
	t.Setenv(
		gosourcedConstant.TokenEnvironment,
		generativeConstant.ModelContextTestToken,
	)
}
