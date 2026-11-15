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

func TestWarningMissingOption(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/r.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.Concern(t, constant.MissingOptionKey, s[0])
}

func TestWarningMissingRun(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/r.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.Concern(t, constant.MissingRunKey, s[0])
}

func TestWarningNoSuffix(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotest/server/r.go", "package server\n")
	v.WriteString("pkg/tool/gotest/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotest/run.go", "package gotest\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	scan_tester.Concern(t, constant.MissingSuffixKey, s[0])
}

func TestWarningRouteExists(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/r.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/route/r.go", "package route\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	scan_tester.Concern(
		t,
		constant.StaleRouteKey,
		scan.Services(v, "test", audit_configuration.New())[0],
	)
}

func TestWarningMissingMountGo(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/server.go",
		"package model_context\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	scan_tester.Concern(
		t,
		constant.MissingMountKey,
		scan.Services(v, "test", audit_configuration.New())[0],
	)
}

func TestWarningMissingCaptureFail(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/mount.go",
		"package model_context\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	scan_tester.Concern(
		t,
		constant.MissingCaptureFailKey,
		scan.Services(v, "test", audit_configuration.New())[0],
	)
}

func TestAnswerMessageSatisfiesCaptureCheck(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/mount.go",
		"package model_context\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/model_context/answer_message.go",
		"package model_context\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	scan_tester.NoConcern(
		t,
		constant.MissingCaptureFailKey,
		scan.Services(v, "test", audit_configuration.New())[0],
	)
}

func TestCleanServiceNoConcerns(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/server/r.go", "package server\n")
	v.WriteString(
		"pkg/tool/gotestd/constant/constant.go",
		"package constant\n\nimport \"github.com/funtimecoding/soil/pkg/identity\"\n\nvar Identity = identity.New(\"gotestd\", \"test\", \"gotestd\")\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", audit_configuration.New())
	assert.Integer(t, 1, len(s))
	assert.Integer(t, 0, len(s[0].Concerns))
}
