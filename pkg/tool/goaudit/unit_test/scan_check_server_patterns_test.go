package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestNilNilReturnFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/generated/server/config.yaml",
		"generate:\n  strict-server: true\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/generated/server/openapi.yaml",
		"---\nopenapi: \"3.0.0\"\ninfo:\n  title: gotestd\n  version: 1.0.0\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/server/get_thing.go",
		"package server\n\nfunc (s *Server) GetThing() (any, error) {\n\treturn nil, nil\n}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], scan.NilNilReturnKey)
}

func TestNilNilReturnClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/generated/server/config.yaml",
		"generate:\n  strict-server: true\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/generated/server/openapi.yaml",
		"---\nopenapi: \"3.0.0\"\ninfo:\n  title: gotestd\n  version: 1.0.0\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/server/get_thing.go",
		"package server\n\nimport \"fmt\"\n\nfunc (s *Server) GetThing() (any, error) {\n\treturn fmt.Sprintf(\"ok\"), nil\n}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], scan.NilNilReturnKey)
}

func TestHttpErrorFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/generated/server/config.yaml",
		"generate:\n  strict-server: true\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/generated/server/openapi.yaml",
		"---\nopenapi: \"3.0.0\"\ninfo:\n  title: gotestd\n  version: 1.0.0\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/server/handler.go",
		"package server\n\nimport \"net/http\"\n\nfunc (s *Server) Handle(w http.ResponseWriter) {\n\thttp.Error(w, \"bad\", 400)\n}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], scan.HttpErrorInStrictKey)
}

func TestHttpErrorClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/generated/server/config.yaml",
		"generate:\n  strict-server: true\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/generated/server/openapi.yaml",
		"---\nopenapi: \"3.0.0\"\ninfo:\n  title: gotestd\n  version: 1.0.0\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/server/handler.go",
		"package server\n\nfunc (s *Server) Handle() string {\n\treturn \"ok\"\n}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], scan.HttpErrorInStrictKey)
}

func TestMissingServerCaptureFailFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/generated/server/config.yaml",
		"generate:\n  strict-server: true\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/generated/server/openapi.yaml",
		"---\npaths:\n  /api/test:\n    get:\n      operationId: getTest\n      responses:\n        \"500\":\n          description: error\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/server/server.go",
		"package server\n\ntype Server struct {\n\treporter any\n}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], scan.MissingServerCaptureFailKey)
}

func TestMissingServerCaptureFailSkipsRecoveryOnly(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/generated/server/config.yaml",
		"generate:\n  strict-server: true\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/generated/server/openapi.yaml",
		"---\npaths:\n  /api/test:\n    get:\n      operationId: getTest\n      responses:\n        \"500\":\n          description: error\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/server/server.go",
		"package server\n\ntype Server struct {\n\tstore any\n}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], scan.MissingServerCaptureFailKey)
}

func assertNoConcern(
	t *testing.T,
	s *scan.Service,
	key string,
) {
	t.Helper()

	for _, c := range s.Concerns {
		if c.Key == key {
			t.Errorf("unexpected concern with key %q found", key)
		}
	}
}
