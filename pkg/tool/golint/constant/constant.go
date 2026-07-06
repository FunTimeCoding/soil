package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"golint",
	"Go source linter with auto-fix support",
	"golint [flags] [path...]",
)
