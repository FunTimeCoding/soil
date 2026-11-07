package option

import "github.com/funtimecoding/soil/pkg/web/option"

type Credential struct {
	option.Web
	Database      string
	Password      string
	RevealedField []string
}
