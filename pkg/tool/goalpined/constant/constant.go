package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goalpined",
	"Alpine package repository server",
	"goalpined",
).WithInstructions(
	"Alpine package repository server - upload, sign, and index apk packages. Use list_packages to see what the repository index holds.",
)

const (
	ListPackages = "list_packages"

	FileAddress = ":8081"
)
