package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gobuild",
	"Cross-compilation and deploy tool",
	"gobuild [command]",
)

const (
	BuildCommand   = "build"
	InstallCommand = "install"
	StaleCommand   = "stale"

	ModuleFlag = "module"
	NativeFlag = "native"
	TagsFlag   = "tags"
	MainFlag   = "main"
	OutputFlag = "output"

	VersionSeparator = "@"
)
