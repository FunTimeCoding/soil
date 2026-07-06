package scan

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"testing"
)

func TestIdentityWarningsForCliTool(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotest/constant/constant.go",
		"package constant\n\nimport \"github.com/funtimecoding/soil/pkg/identity\"\n\nvar Identity = identity.New(\"wrong\", \"test\", \"wrong\")\n",
	)
	w := IdentityWarnings(v, nil)
	assert.Integer(t, 1, len(w))
	assert.String(t, IdentityMismatchKey, w[0].Key)
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
		len(IdentityWarnings(v, Services(v, "test", &Configuration{}))),
	)
}

func TestIdentityWarningsSkipsNoConstant(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotest/main.go", "package gotest\n")
	assert.Integer(t, 0, len(IdentityWarnings(v, nil)))
}
