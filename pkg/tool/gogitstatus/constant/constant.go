package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gogitstatus",
	"Git repository status check",
	"gogitstatus [flags]",
)
