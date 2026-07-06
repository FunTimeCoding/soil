package scan

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"testing"
)

func TestWarningMissingOption(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/server/r.go",
		"package server\n",
	)
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := Services(v, "test", &Configuration{})
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], MissingOptionKey)
}

func TestWarningMissingRun(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/server/r.go",
		"package server\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/option/o.go",
		"package option\n",
	)
	s := Services(v, "test", &Configuration{})
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], MissingRunKey)
}

func TestWarningNoSuffix(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotest/server/r.go",
		"package server\n",
	)
	v.WriteString(
		"pkg/tool/gotest/option/o.go",
		"package option\n",
	)
	v.WriteString("pkg/tool/gotest/run.go", "package gotest\n")
	s := Services(v, "test", &Configuration{})
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], MissingSuffixKey)
}

func TestWarningRouteExists(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/server/r.go",
		"package server\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/route/r.go",
		"package route\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/option/o.go",
		"package option\n",
	)
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	assertConcern(t, Services(v, "test", &Configuration{})[0], StaleRouteKey)
}

func TestWarningMissingMountGo(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/server.go",
		"package model_context\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/option/o.go",
		"package option\n",
	)
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	assertConcern(
		t,
		Services(v, "test", &Configuration{})[0],
		MissingMountKey,
	)
}

func TestWarningMissingCaptureFail(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/mount.go",
		"package model_context\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/option/o.go",
		"package option\n",
	)
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	assertConcern(
		t,
		Services(v, "test", &Configuration{})[0],
		MissingCaptureFailKey,
	)
}

func TestWarningConstantFile(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/server/r.go",
		"package server\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/constant.go",
		"package gotestd\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/option/o.go",
		"package option\n",
	)
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	assertConcern(
		t,
		Services(v, "test", &Configuration{})[0],
		ConstantFileKey,
	)
}

func TestCleanServiceNoConcerns(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/server/r.go",
		"package server\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/constant/constant.go",
		"package constant\n\nimport \"github.com/funtimecoding/soil/pkg/identity\"\n\nvar Identity = identity.New(\"gotestd\", \"test\", \"gotestd\")\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/option/o.go",
		"package option\n",
	)
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := Services(v, "test", &Configuration{})
	assert.Integer(t, 1, len(s))
	assert.Integer(t, 0, len(s[0].Concerns))
}

func assertConcern(
	t *testing.T,
	s *Service,
	key string,
) {
	t.Helper()

	for _, c := range s.Concerns {
		if c.Key == key {
			return
		}
	}

	t.Errorf("expected concern with key %q not found", key)
}
