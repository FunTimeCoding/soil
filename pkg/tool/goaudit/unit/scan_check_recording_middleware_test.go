package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/unit/scan_tester"
	"testing"
)

func TestRecordingMiddlewareNilFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/s.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/mount.go",
		`package gotestd

func Mount(g *guard.Mux) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(server.New(s, r), nil),
			http.NewServeMux(),
		),
	)
}
`,
	)
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.Concern(t, constant.StrictMiddlewareKey, s[0])
}

func TestRecordingMiddlewareDaemonMountFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/s.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/daemon/mount.go",
		`package daemon

func (d *Daemon) Mount(g *guard.Mux) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(server.New(s, r), nil),
			http.NewServeMux(),
		),
	)
}
`,
	)
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.Concern(t, constant.StrictMiddlewareKey, s[0])
}

func TestRecordingMiddlewareInlineFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/s.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/mount.go",
		`package gotestd

func Mount(g *guard.Mux) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(s, r),
				[]generated.StrictMiddlewareFunc{
					func(
						f generated.StrictHandlerFunc,
						operation string,
					) generated.StrictHandlerFunc {
						return f
					},
				},
			),
			http.NewServeMux(),
		),
	)
}
`,
	)
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.Concern(t, constant.StrictMiddlewareKey, s[0])
}

func TestRecordingMiddlewareClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/s.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/mount.go",
		`package gotestd

func Mount(g *guard.Mux) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(s, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
}
`,
	)
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.NoConcern(t, constant.StrictMiddlewareKey, s[0])
}
