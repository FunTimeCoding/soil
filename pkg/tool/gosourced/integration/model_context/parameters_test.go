package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/model_context_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestAddParametersThroughTheTool(t *testing.T) {
	o := model_context_tester.New(
		t,
		"../../service/testdata/add-parameters/src",
	)
	o.Client.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := o.Client.MustCallTool(
		constant.AddParameters,
		map[string]any{
			"functions": []any{
				map[string]any{
					"package_path": "example/pkg/target",
					"name":         "Answer",
					"parameter":    "x",
					"type":         "context.Context",
					"position":     0,
					"argument":     "x",
				},
			},
		},
	)
	assert.StringContains(t, "Answer: passed x", result)
	caller, e := os.ReadFile(filepath.Join(o.Directory, "pkg/caller/caller.go"))
	assert.FatalOnError(t, e)
	assert.StringContains(t, `target.Answer(x, "a", 1)`, string(caller))
}

func TestNameParametersThroughTheTool(t *testing.T) {
	o := model_context_tester.New(
		t,
		"../../service/testdata/name-parameters/src",
	)
	o.Client.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := o.Client.MustCallTool(
		constant.NameParameters,
		map[string]any{
			"package_path": "example/pkg/handler",
			"functions": []any{
				map[string]any{"name": "Get", "receiver": "Server"},
			},
			"type": "context.Context",
			"name": "x",
		},
	)
	assert.StringContains(t, "Get: named x", result)
	source, e := os.ReadFile(
		filepath.Join(o.Directory, "pkg/handler/handler.go"),
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "\tx context.Context,\n", string(source))
}

func TestAddParametersToolRefusesWithoutAnArgument(t *testing.T) {
	o := model_context_tester.New(
		t,
		"../../service/testdata/add-parameters/src",
	)
	o.Client.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	_, e := o.Client.CallTool(
		constant.AddParameters,
		map[string]any{
			"functions": []any{
				map[string]any{
					"package_path": "example/pkg/target",
					"name":         "Answer",
					"parameter":    "x",
					"type":         "context.Context",
					"position":     0,
					"argument":     "",
				},
			},
		},
	)
	assert.StringContains(t, "needs parameter, type and argument", e.Error())
}
