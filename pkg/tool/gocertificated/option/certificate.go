package option

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Certificate struct {
	option.Web
	PostgresLocator    string
	LitePath           string
	Project            string
	Branch             string
	AuthorityDirectory string
	SecretAuthority    string
	SecretPath         string
	Forge              func() face.GitLabClient
}
