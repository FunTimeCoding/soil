package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"golicense",
	"Dependency license check",
	"golicense [flags] [path]",
)
