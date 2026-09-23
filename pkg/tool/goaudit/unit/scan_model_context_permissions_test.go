package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/unit/scan_tester"
	"testing"
)

func TestModelContextPermissionsClean(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteAlfaService(
		base,
		`{"permissions":{"allow":["mcp__alfa__alfa_search","Bash(ls:*)"]}}`,
	)
	result := scan.ModelContextPermissions(
		base,
		scan_tester.AlfaConfiguration(),
	)
	assert.Integer(t, 0, len(result))
}

func TestModelContextPermissionsStale(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteAlfaService(
		base,
		`{"permissions":{"allow":["mcp__alfa__alfa_search","mcp__alfa__old_name"]}}`,
	)
	result := scan.ModelContextPermissions(
		base,
		scan_tester.AlfaConfiguration(),
	)
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
	scan_tester.WriteAlfaService(
		base,
		`{"permissions":{"deny":["mcp__alfa__gone"],"ask":["mcp__alfa__also_gone"]}}`,
	)
	result := scan.ModelContextPermissions(
		base,
		scan_tester.AlfaConfiguration(),
	)
	assert.Integer(t, 2, len(result))
}

func TestModelContextPermissionsUnmappedServerSkipped(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteAlfaService(
		base,
		`{"permissions":{"allow":["mcp__github__search_code"]}}`,
	)
	result := scan.ModelContextPermissions(
		base,
		scan_tester.AlfaConfiguration(),
	)
	assert.Integer(t, 0, len(result))
}

func TestModelContextPermissionsServerEntrySkipped(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteAlfaService(
		base,
		`{"permissions":{"allow":["mcp__alfa"]}}`,
	)
	result := scan.ModelContextPermissions(
		base,
		scan_tester.AlfaConfiguration(),
	)
	assert.Integer(t, 0, len(result))
}

func TestModelContextPermissionsMappedPathMissing(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteAlfaService(
		base,
		`{"permissions":{"allow":["mcp__bravo__anything"]}}`,
	)
	configuration := scan.NewConfiguration()
	configuration.ModelContext = map[string]string{"bravo": "pkg/tool/gobravod"}
	result := scan.ModelContextPermissions(base, configuration)
	assert.Integer(t, 1, len(result))
	assert.String(t, "mapped_service_missing", result[0].Key)
	assert.String(t, "pkg/tool/gobravod", result[0].Path)
}

func TestModelContextPermissionsNoSettings(t *testing.T) {
	configuration := scan_tester.AlfaConfiguration()
	result := scan.ModelContextPermissions(t.TempDir(), configuration)
	assert.Integer(t, 0, len(result))
}

func TestModelContextPermissionsNoMap(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteAlfaService(
		base,
		`{"permissions":{"allow":["mcp__alfa__alfa_search"]}}`,
	)
	result := scan.ModelContextPermissions(base, scan.NewConfiguration())
	assert.Integer(t, 0, len(result))
}
