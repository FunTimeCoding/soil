package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"regexp"
)

var Identity = identity.New(
	"goaudit",
	"Compliance matrix for services and clients",
	"goaudit <repo-root> [<repo-root>...]",
)

var (
	ConfigurationPaths = []string{".goaudit.yaml", "strata/tool/goaudit.yaml"}
	VersionedPathPattern = regexp.MustCompile(`^/api/v[0-9]+(/|$)`)
)

const (
	ClaudeSettingsPath       = ".claude/settings.local.json"
	ConstantDirectory        = "constant"
	ConstantFileName         = "constant.go"
	IntegrationDirectory     = "integration"
	ModelContextDirectory    = "model_context"
	PackageDirectory         = "pkg"
	TestdataDirectory        = "testdata"
	ToolDirectory            = "tool"
	UnitDirectory            = "unit"
)
