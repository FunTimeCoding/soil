package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestIdentityWarningsForCliTool(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotest/constant/constant.go",
		"package constant\n\nimport \"github.com/funtimecoding/soil/pkg/identity\"\n\nvar Identity = identity.New(\"wrong\", \"test\", \"wrong\")\n",
	)
	w := scan.IdentityWarnings(v, nil)
	assert.Integer(t, 1, len(w))
	assert.String(t, scan.IdentityMismatchKey, w[0].Key)
}

func TestIdentityWarningsSkipsServices(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/r.go", "package server\n")
	v.WriteString(
		"pkg/tool/gotestd/constant/constant.go",
		"package constant\n\nimport \"github.com/funtimecoding/soil/pkg/identity\"\n\nvar Identity = identity.New(\"wrong\", \"test\", \"wrong\")\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	assert.Integer(
		t,
		0,
		len(
			scan.IdentityWarnings(
				v,
				scan.Services(v, "test", scan.NewConfiguration()),
			),
		),
	)
}

func TestIdentityWarningsSkipsNoConstant(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotest/main.go", "package gotest\n")
	assert.Integer(t, 0, len(scan.IdentityWarnings(v, nil)))
}
