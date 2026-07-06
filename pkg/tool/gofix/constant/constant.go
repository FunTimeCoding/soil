package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gofix",
	"Go source fixer",
	"gofix [flags]",
)

const MaxSingleParameterLength = 80
