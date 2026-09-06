package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestRawModelContextServerFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		`package model_context

func New(r face.Reporter) *Server {
	return &Server{
		server: server.NewMCPServer(
			"test",
			constant.DefaultVersion,
			server.WithToolCapabilities(true),
		),
	}
}
`,
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], constant.RawModelContextServerKey)
}

func TestRawModelContextServerClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		`package model_context

func New(
	r face.Reporter,
	t face.Recorder,
	version string,
) *Server {
	return &Server{
		server: server.New(
			constant.Identity,
			version,
		).WithRecorder(t).Server(),
	}
}
`,
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], constant.RawModelContextServerKey)
}
