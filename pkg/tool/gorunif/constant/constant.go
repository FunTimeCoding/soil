package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gorunif",
	"Conditional command runner",
	"gorunif [flags] <command>",
)

const (
	Base   = "base"
	Head   = "head"
	Suffix = "suffix"
)
