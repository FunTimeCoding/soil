package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestVersionedPathFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/generated/server/openapi.yaml",
		"---\nopenapi: \"3.0.0\"\ninfo:\n  title: gotestd\n  version: 1.0.0\npaths:\n  /api/v1/things:\n    get: {}\n  /api/status:\n    get: {}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], constant.VersionedPathKey)
}

func TestVersionedPathClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/generated/server/openapi.yaml",
		"---\nopenapi: \"3.0.0\"\ninfo:\n  title: gotestd\n  version: 1.0.0\npaths:\n  /api/things:\n    get: {}\n  /api/views:\n    get: {}\n  /api/versions:\n    get: {}\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], constant.VersionedPathKey)
}
