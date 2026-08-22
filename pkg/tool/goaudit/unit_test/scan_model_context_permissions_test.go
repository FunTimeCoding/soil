package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"path/filepath"
	"testing"
)

func writeAlfaService(base string, settings string) {
	system.MakeDirectory(filepath.Join(base, ".claude"))
	system.WriteFile(
		filepath.Join(base, ".claude", "settings.local.json"),
		[]byte(settings),
		0644,
	)
	system.MakeDirectory(filepath.Join(base, "pkg/tool/goalfad/constant"))
	system.MakeDirectory(filepath.Join(base, "pkg/tool/goalfad/model_context"))
	system.WriteFile(
		filepath.Join(base, "pkg/tool/goalfad/constant/constant.go"),
		[]byte("package constant\n\nconst AlfaSearch = \"alfa_search\"\n"),
		0644,
	)
	system.WriteFile(
		filepath.Join(base, "pkg/tool/goalfad/model_context/register.go"),
		[]byte(
			"package model_context\n\nfunc register() {\n\tadd(mcp.NewTool(constant.AlfaSearch))\n}\n",
		),
		0644,
	)
}

func alfaConfiguration() *scan.Configuration {
	configuration := scan.NewConfiguration()
	configuration.ModelContext = map[string]string{"alfa": "pkg/tool/goalfad"}

	return configuration
}

func TestModelContextPermissionsClean(t *testing.T) {
	base := t.TempDir()
	writeAlfaService(
		base,
		`{"permissions":{"allow":["mcp__alfa__alfa_search","Bash(ls:*)"]}}`,
	)
	result := scan.ModelContextPermissions(base, alfaConfiguration())
	assert.Integer(t, 0, len(result))
}

func TestModelContextPermissionsStale(t *testing.T) {
	base := t.TempDir()
	writeAlfaService(
		base,
		`{"permissions":{"allow":["mcp__alfa__alfa_search","mcp__alfa__old_name"]}}`,
	)
	result := scan.ModelContextPermissions(base, alfaConfiguration())
	assert.Integer(t, 1, len(result))
	assert.String(t, "stale_tool_permission", result[0].Key)
	assert.String(t, ".claude/settings.local.json", result[0].Path)
	assert.String(
		t,
		"mcp__alfa__old_name not registered by pkg/tool/goalfad",
		result[0].Text,
	)
}

func TestModelContextPermissionsDenyAndAsk(t *testing.T) {
	base := t.TempDir()
	writeAlfaService(
		base,
		`{"permissions":{"deny":["mcp__alfa__gone"],"ask":["mcp__alfa__also_gone"]}}`,
	)
	result := scan.ModelContextPermissions(base, alfaConfiguration())
	assert.Integer(t, 2, len(result))
}

func TestModelContextPermissionsUnmappedServerSkipped(t *testing.T) {
	base := t.TempDir()
	writeAlfaService(
		base,
		`{"permissions":{"allow":["mcp__github__search_code"]}}`,
	)
	result := scan.ModelContextPermissions(base, alfaConfiguration())
	assert.Integer(t, 0, len(result))
}

func TestModelContextPermissionsServerEntrySkipped(t *testing.T) {
	base := t.TempDir()
	writeAlfaService(base, `{"permissions":{"allow":["mcp__alfa"]}}`)
	result := scan.ModelContextPermissions(base, alfaConfiguration())
	assert.Integer(t, 0, len(result))
}

func TestModelContextPermissionsMappedPathMissing(t *testing.T) {
	base := t.TempDir()
	writeAlfaService(base, `{"permissions":{"allow":["mcp__bravo__anything"]}}`)
	configuration := scan.NewConfiguration()
	configuration.ModelContext = map[string]string{"bravo": "pkg/tool/gobravod"}
	result := scan.ModelContextPermissions(base, configuration)
	assert.Integer(t, 1, len(result))
	assert.String(t, "mapped_service_missing", result[0].Key)
	assert.String(t, "pkg/tool/gobravod", result[0].Path)
}

func TestModelContextPermissionsNoSettings(t *testing.T) {
	configuration := alfaConfiguration()
	result := scan.ModelContextPermissions(t.TempDir(), configuration)
	assert.Integer(t, 0, len(result))
}

func TestModelContextPermissionsNoMap(t *testing.T) {
	base := t.TempDir()
	writeAlfaService(
		base,
		`{"permissions":{"allow":["mcp__alfa__alfa_search"]}}`,
	)
	result := scan.ModelContextPermissions(base, scan.NewConfiguration())
	assert.Integer(t, 0, len(result))
}
