package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"path/filepath"
	"testing"
)

func TestLoadToolPath(t *testing.T) {
	root := t.TempDir()
	write(
		root,
		constant.ToolFile,
		"pre-push:\n  - paths: ['**/*.go']\n    run: task lint\n  - run: task always\n",
	)
	c := configuration.Load(root)
	assert.String(t, filepath.Join(root, "strata/tool/gohook.yaml"), c.Path)
	assert.Strings(t, []string{"pre-push"}, c.HookNames())
	assert.Count(t, 2, c.Hooks["pre-push"])
	assert.String(t, "task lint", c.Hooks["pre-push"][0].Run)
	assert.Strings(t, []string{"**/*.go"}, c.Hooks["pre-push"][0].Paths)
	assert.Count(t, 0, c.Hooks["pre-push"][1].Paths)
}

func TestLoadRootBeatsTool(t *testing.T) {
	root := t.TempDir()
	write(root, constant.RootFile, "pre-push:\n  - run: root\n")
	write(root, constant.ToolFile, "pre-push:\n  - run: tool\n")
	c := configuration.Load(root)
	assert.String(t, filepath.Join(root, ".gohook.yaml"), c.Path)
	assert.String(t, "root", c.Hooks["pre-push"][0].Run)
}

func TestLoadMissingPanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	configuration.Load(t.TempDir())
}

func TestLoadUnknownHookPanics(t *testing.T) {
	root := t.TempDir()
	write(root, constant.RootFile, "pre-psuh:\n  - run: x\n")
	defer func() { assert.NotNil(t, recover()) }()
	configuration.Load(root)
}

func TestLoadMissingRunPanics(t *testing.T) {
	root := t.TempDir()
	write(root, constant.RootFile, "pre-push:\n  - paths: [a/]\n")
	defer func() { assert.NotNil(t, recover()) }()
	configuration.Load(root)
}

func TestHookNamesSorted(t *testing.T) {
	root := t.TempDir()
	write(
		root,
		constant.RootFile,
		"pre-push:\n  - run: a\npre-commit:\n  - run: b\n",
	)
	assert.Strings(
		t,
		[]string{"pre-commit", "pre-push"},
		configuration.Load(root).HookNames(),
	)
}
