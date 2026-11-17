package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"os"
	"path/filepath"
	"testing"
)

func TestAddParametersThroughTool(t *testing.T) {
	s := base.New(
		t,
		base.WithFixture("../../service/testdata/add-parameters/src"),
	)
	c := s.Served().Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := c.MustCallTool(
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
	caller, e := os.ReadFile(filepath.Join(s.Directory, "pkg/caller/caller.go"))
	assert.FatalOnError(t, e)
	assert.StringContains(t, `target.Answer(x, "a", 1)`, string(caller))
}

func TestNameParametersThroughTool(t *testing.T) {
	s := base.New(
		t,
		base.WithFixture("../../service/testdata/name-parameters/src"),
	)
	c := s.Served().Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := c.MustCallTool(
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
		filepath.Join(s.Directory, "pkg/handler/handler.go"),
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "\tx context.Context,\n", string(source))
}

func TestAddParametersToolRefusesWithoutArgument(t *testing.T) {
	s := base.New(
		t,
		base.WithFixture("../../service/testdata/add-parameters/src"),
	)
	c := s.Served().Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	_, e := c.CallTool(
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
