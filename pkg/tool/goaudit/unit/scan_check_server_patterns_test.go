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
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).Concern(constant.NilNilReturnKey, s[0])
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
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).NoConcern(constant.NilNilReturnKey, s[0])
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
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).Concern(constant.HttpErrorInStrictKey, s[0])
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
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).NoConcern(constant.HttpErrorInStrictKey, s[0])
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
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).Concern(constant.MissingServerCaptureFailKey, s[0])
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
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).NoConcern(constant.MissingServerCaptureFailKey, s[0])
}

func TestMissingCaptureDetailFlagged(t *testing.T) {
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
		"package server\n\nimport \"github.com/funtimecoding/soil/pkg/web/detail_error\"\n\nfunc (s *Server) GetThing() *detail_error.Detail {\n\treturn nil\n}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).Concern(constant.MissingCaptureDetailKey, s[0])
}

func TestMissingCaptureDetailClean(t *testing.T) {
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
		"package server\n\nimport \"github.com/funtimecoding/soil/pkg/web/detail_error\"\n\nfunc (s *Server) GetThing() *detail_error.Detail {\n\treturn nil\n}\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/server/capture_detail.go",
		"package server\n\nfunc (s *Server) captureDetail() {}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).NoConcern(constant.MissingCaptureDetailKey, s[0])
}

func TestMissingCaptureDetailNotWrapping(t *testing.T) {
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
		"package server\n\nfunc (s *Server) GetThing() string {\n\treturn \"ok\"\n}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).NoConcern(constant.MissingCaptureDetailKey, s[0])
}
