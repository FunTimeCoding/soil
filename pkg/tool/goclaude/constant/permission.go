package constant

import "regexp"

const (
	ConfigurationDirectoryEnvironment = "CLAUDE_CONFIG_DIR"
	ProjectDirectoryEnvironment       = "CLAUDE_PROJECT_DIR"
	ClaudeDirectory                   = ".claude"
	LabelAllowedWarning               = "[goclauded] the label tool is auto-allowed in %s - a session could grant itself git-write; remove it from permissions.allow"
)

var (
	SettingsFiles        = []string{"settings.json", "settings.local.json"}
	ManagedSettingsFiles = []string{
		"/Library/Application Support/ClaudeCode/managed-settings.json",
		"/etc/claude-code/managed-settings.json",
	}
	LabelToolAllowance = regexp.MustCompile(
		`^mcp__(.+_)?goclauded(__(label|\*))?$`,
	)
)
