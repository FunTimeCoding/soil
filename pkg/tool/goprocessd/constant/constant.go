package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goprocessd",
	"Process manager with environment reload",
	"goprocessd [flags]",
)

const HistoryCapacity = 200
