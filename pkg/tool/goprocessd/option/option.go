package option

import "github.com/funtimecoding/soil/pkg/web/option"

type Option struct {
	option.Web
	Local        option.Web
	ProcfilePath string
	EnvrcPath    string
}
