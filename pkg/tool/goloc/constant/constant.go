package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goloc",
	"Count code, comment and blank lines by language",
	"goloc [flags] [path...]",
)
