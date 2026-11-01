package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goaptly",
	"Debian package repository CLI",
	"goaptly [command]",
)

const (
	RepositoryFlag     = "repository"
	ForceFlag          = "force"
	DistributionFlag   = "distribution"
	PassphraseFileFlag = "passphrase-file"

	DefaultRepository = "rolling"
)
