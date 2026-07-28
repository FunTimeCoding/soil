package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goaudit",
	"Compliance matrix for services and clients",
	"goaudit <repo-root> [<repo-root>...]",
)

const (
	ConstantDirectory        = "constant"
	ConstantFileName         = "constant.go"
	IntegrationTestDirectory = "integration_test"
	PackageDirectory         = "pkg"
	TestdataDirectory        = "testdata"
	ToolDirectory            = "tool"
	UnitTestDirectory        = "unit_test"
)
