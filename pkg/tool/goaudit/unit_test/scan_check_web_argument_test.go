package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestWebArgumentPortFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/s.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/main.go",
		`package gotestd

func Main() {
	a := argument.NewInstance(constant.Identity)
	a.Integer(argumentConstant.Port, web.ListenPort, web.PortUsage)
	a.Parse(version, gitHash, buildDate)
}
`,
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], constant.WebArgumentKey)
}

func TestWebArgumentMetricPortFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/s.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/main.go",
		`package gotestd

func Main() {
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Integer(argumentConstant.MetricPort, metric.Port, metric.PortUsage)
	a.Parse(version, gitHash, buildDate)
}
`,
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], constant.WebArgumentKey)
}

func TestWebArgumentClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/s.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/main.go",
		`package gotestd

func Main() {
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Metric()
	a.Integer(argumentConstant.Interval, 5, "Poll interval")
	a.Parse(version, gitHash, buildDate)
}
`,
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], constant.WebArgumentKey)
}
