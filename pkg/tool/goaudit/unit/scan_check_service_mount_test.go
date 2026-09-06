package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
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
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], constant.MissingServiceMountKey)
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
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], constant.MissingServiceMountKey)
}
