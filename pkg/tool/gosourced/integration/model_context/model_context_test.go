package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"os"
	"path/filepath"
	"testing"
)

func TestListModules(t *testing.T) {
	s := base.New(t, "../../service/testdata/unexport-function/src")
	c := s.Connect(t)
	result := c.MustCallTool(constant.ListModules, map[string]any{})
	assert.StringContains(t, "test", result)
}

func TestUseModule(t *testing.T) {
	s := base.New(t, "../../service/testdata/unexport-function/src")
	c := s.Connect(t)
	result := c.MustCallTool(
		constant.UseModule,
		map[string]any{"module": "test"},
	)
	assert.StringContains(t, "active module set to test", result)
}

func TestUseModuleUnknown(t *testing.T) {
	s := base.New(t, "../../service/testdata/unexport-function/src")
	c := s.Connect(t)
	_, e := c.CallTool(constant.UseModule, map[string]any{"module": "missing"})
	assert.StringContains(t, "unknown module", e.Error())
}

func TestChangeVisibilityExport(t *testing.T) {
	s := base.New(t, "../../service/testdata/export-function/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := c.MustCallTool(
		constant.ChangeVisibility,
		map[string]any{
			"symbol":       "isGenerated",
			"package_path": "example/pkg/target",
		},
	)
	assert.StringContains(t, "isGenerated → IsGenerated", result)
	helper, e := os.ReadFile(
		filepath.Join(s.Directory, "pkg/target/is_generated.go"),
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "func IsGenerated(", string(helper))
}

func TestChangeVisibilityNoModule(t *testing.T) {
	s := base.New(t, "../../service/testdata/unexport-function/src")
	c := s.Connect(t)
	_, e := c.CallTool(
		constant.ChangeVisibility,
		map[string]any{
			"symbol":       "IsGenerated",
			"package_path": "example/pkg/target",
		},
	)
	assert.StringContains(t, "use_module", e.Error())
}

func TestChangeVisibilityCrossPackageBlocked(t *testing.T) {
	s := base.New(t, "../../service/testdata/cross-package/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	_, e := c.CallTool(
		constant.ChangeVisibility,
		map[string]any{
			"symbol":       "IsValid",
			"package_path": "example/pkg/target",
		},
	)
	assert.StringContains(t, "would lose access", e.Error())
}

func TestRenameSymbol(t *testing.T) {
	s := base.New(t, "../../service/testdata/rename-function/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := c.MustCallTool(
		constant.RenameSymbol,
		map[string]any{
			"package_path": "example/pkg/target",
			"old_name":     "IsGeneratedHeader",
			"new_name":     "IsGenerated",
		},
	)
	assert.StringContains(t, "IsGeneratedHeader → IsGenerated", result)
	helper, e := os.ReadFile(
		filepath.Join(s.Directory, "pkg/target/is_generated.go"),
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "func IsGenerated(", string(helper))
}

func TestRenameSymbolNamesPackageThatDoesNotLoad(t *testing.T) {
	s := base.New(t, "../../service/testdata/rename-function/src")
	c := s.Connect(t)
	system_tester.Write(
		t,
		s.Directory,
		"pkg/broken/broken.go",
		"package broken\n\nimport \"example/pkg/target\"\n\nfunc Broken() bool {\n\treturn target.IsGeneratedHeader(missing)\n}\n",
	)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	_, e := c.CallTool(
		constant.RenameSymbol,
		map[string]any{
			"package_path": "example/pkg/target",
			"old_name":     "IsGeneratedHeader",
			"new_name":     "IsGenerated",
		},
	)
	assert.StringContains(
		t,
		"module does not load: example/pkg/broken",
		e.Error(),
	)
}

func TestRenameSymbolToUnexportedBlocked(t *testing.T) {
	s := base.New(t, "../../service/testdata/rename-unexport/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	_, e := c.CallTool(
		constant.RenameSymbol,
		map[string]any{
			"package_path": "example/pkg/target",
			"old_name":     "Validate",
			"new_name":     "check",
		},
	)
	assert.StringContains(t, "would lose access", e.Error())
}

func TestRenameSymbolSameName(t *testing.T) {
	s := base.New(t, "../../service/testdata/rename-function/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	_, e := c.CallTool(
		constant.RenameSymbol,
		map[string]any{
			"package_path": "example/pkg/target",
			"old_name":     "IsGeneratedHeader",
			"new_name":     "IsGeneratedHeader",
		},
	)
	assert.StringContains(t, "old_name and new_name are the same", e.Error())
}

func TestRenameSymbolNoModule(t *testing.T) {
	s := base.New(t, "../../service/testdata/rename-function/src")
	c := s.Connect(t)
	_, e := c.CallTool(
		constant.RenameSymbol,
		map[string]any{
			"package_path": "example/pkg/target",
			"old_name":     "IsGeneratedHeader",
			"new_name":     "IsGenerated",
		},
	)
	assert.StringContains(t, "use_module", e.Error())
}

func TestRenameSymbolMissingParams(t *testing.T) {
	s := base.New(t, "../../service/testdata/rename-function/src")
	c := s.Connect(t)
	_, e := c.CallTool(
		constant.RenameSymbol,
		map[string]any{
			"old_name": "IsGeneratedHeader",
			"new_name": "IsGenerated",
		},
	)
	assert.StringContains(t, "input schema validation failed", e.Error())
	assert.StringContains(t, "missing property 'package_path'", e.Error())
}

func TestRenameSymbolNotFound(t *testing.T) {
	s := base.New(t, "../../service/testdata/rename-function/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	_, e := c.CallTool(
		constant.RenameSymbol,
		map[string]any{
			"package_path": "example/pkg/target",
			"old_name":     "Missing",
			"new_name":     "Something",
		},
	)
	assert.StringContains(t, "not found", e.Error())
}

func TestExtractToFile(t *testing.T) {
	s := base.New(t, "../../service/testdata/extract-function/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := c.MustCallTool(
		constant.ExtractToFile,
		map[string]any{
			"file":   "pkg/target/combined.go",
			"symbol": "FormatName",
		},
	)
	assert.StringContains(t, "format_name.go", result)
}

func TestExtractToFileNotFound(t *testing.T) {
	s := base.New(t, "../../service/testdata/extract-function/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	_, e := c.CallTool(
		constant.ExtractToFile,
		map[string]any{"file": "pkg/target/combined.go", "symbol": "Missing"},
	)
	assert.StringContains(t, "not found", e.Error())
}

func TestExtractToFileNoModule(t *testing.T) {
	s := base.New(t, "../../service/testdata/extract-function/src")
	c := s.Connect(t)
	_, e := c.CallTool(
		constant.ExtractToFile,
		map[string]any{
			"file":   "pkg/target/combined.go",
			"symbol": "FormatName",
		},
	)
	assert.StringContains(t, "use_module", e.Error())
}

func TestExtractToFileRenamesSource(t *testing.T) {
	s := base.New(t, "../../service/testdata/extract-last-pair/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := c.MustCallTool(
		constant.ExtractToFile,
		map[string]any{
			"file":   "pkg/target/combined.go",
			"symbol": "FormatName",
		},
	)
	assert.StringContains(t, "format_name.go", result)
	assert.StringContains(t, "summarize_name.go", result)
	_, e := os.Stat(filepath.Join(s.Directory, "pkg/target/combined.go"))
	assert.True(t, os.IsNotExist(e))
}

func TestAddImport(t *testing.T) {
	s := base.New(t, "../../service/testdata/import-empty/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := c.MustCallTool(
		constant.AddImport,
		map[string]any{
			"file":        "pkg/target/example.go",
			"import_path": "fmt",
		},
	)
	assert.StringContains(t, "added fmt", result)
}

func TestRemoveImport(t *testing.T) {
	s := base.New(t, "../../service/testdata/import-grouped/src")
	c := s.Connect(t)
	c.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := c.MustCallTool(
		constant.RemoveImport,
		map[string]any{
			"file":        "pkg/target/example.go",
			"import_path": "strings",
		},
	)
	assert.StringContains(t, "removed strings", result)
}
