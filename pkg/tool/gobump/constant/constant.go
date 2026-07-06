package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gobump",
	"Semantic version bumper",
	"gobump [flags]",
)
