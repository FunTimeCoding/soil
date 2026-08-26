package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goalpine",
	"Alpine package repository CLI",
	"goalpine [command]",
)

const (
	VersionFlag      = "version"
	RepositoryFlag   = "repository"
	ArchitectureFlag = "architecture"

	DefaultVersion    = "rolling"
	DefaultRepository = "main"
)
