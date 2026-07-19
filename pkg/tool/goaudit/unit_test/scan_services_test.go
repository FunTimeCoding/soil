package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestServicesDetectsCapabilities(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/router.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/option.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	services := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(services))
	assert.String(t, "gotestd", services[0].Name)
	assert.Boolean(t, true, services[0].Server)
}

func TestServicesSkipsNonGoPrefix(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/other/server/router.go", "package server\n")
	services := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 0, len(services))
}

func TestServicesSkipsNoCapabilities(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotest/main.go", "package gotest\n")
	services := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 0, len(services))
}

func TestServicesSorted(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gobravod/server/r.go", "package server\n")
	v.WriteString("pkg/tool/gobravod/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gobravod/run.go", "package gobravod\n")
	v.WriteString("pkg/tool/goalfad/server/r.go", "package server\n")
	v.WriteString("pkg/tool/goalfad/option/o.go", "package option\n")
	v.WriteString("pkg/tool/goalfad/run.go", "package goalfad\n")
	services := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 2, len(services))
	assert.String(t, "goalfad", services[0].Name)
	assert.String(t, "gobravod", services[1].Name)
}
