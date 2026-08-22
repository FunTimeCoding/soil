package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goaudit",
	"Compliance matrix for services and clients",
	"goaudit <repo-root> [<repo-root>...]",
)

var ConfigurationPaths = []string{".goaudit.yaml", "strata/tool/goaudit.yaml"}

const (
	ClaudeSettingsPath       = ".claude/settings.local.json"
	ConstantDirectory        = "constant"
	ConstantFileName         = "constant.go"
	IntegrationTestDirectory = "integration_test"
	ModelContextDirectory    = "model_context"
	PackageDirectory         = "pkg"
	TestdataDirectory        = "testdata"
	ToolDirectory            = "tool"
	UnitTestDirectory        = "unit_test"
)
