package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func rootSpec() string {
	return `info:
  title: Test
paths:
  /metrics:
    get: {}
  /raw:
    get: {}
  /api/power:
    get: {}
`
}

func TestRootRouteUncoveredFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString("pkg/tool/gotestd/generated/server/openapi.yaml", rootSpec())
	v.WriteString(
		"pkg/tool/gotestd/mount.go",
		`package gotestd

func Mount(g *guard.Mux) {
	g.TokenMount(constant.InterfacePath, h)
}
`,
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], constant.RootRouteKey)
}

func TestRootRouteExplicitMountsClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString("pkg/tool/gotestd/generated/server/openapi.yaml", rootSpec())
	v.WriteString(
		"pkg/tool/gotestd/constant/constant.go",
		`package constant

const RawPath = "/raw"
`,
	)
	v.WriteString(
		"pkg/tool/gotestd/mount.go",
		`package gotestd

func Mount(g *guard.Mux) {
	g.TokenMount(constant.InterfacePath, h)
	g.OpenMount(route.Get(constant.MetricsPath), h)
	g.TokenMount(route.Get(testConstant.RawPath), h)
}
`,
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], constant.RootRouteKey)
}

func TestRootRouteSlashMountClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString("pkg/tool/gotestd/generated/server/openapi.yaml", rootSpec())
	v.WriteString(
		"pkg/tool/gotestd/mount.go",
		`package gotestd

func Mount(g *guard.Mux) {
	g.TokenMount(constant.Slash, h)
}
`,
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], constant.RootRouteKey)
}
