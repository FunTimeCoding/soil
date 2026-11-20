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

func TestServiceMountMissingFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).Concern(constant.MissingServiceMountKey, s[0])
}

func TestServiceMountPresentClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString("pkg/tool/gotestd/mount.go", "package gotestd\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.New(t).NoConcern(constant.MissingServiceMountKey, s[0])
}
