package option

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Option struct {
	option.Web
	PostgresLocator       string
	LitePath              string
	Secret                string
	SuperUserMail         string
	SuperUserPassword     string
	Issuer                string
	AdminClientIdentifier string
	AdminClientSecret     string
	Directory             func() face.DirectoryClient
}
