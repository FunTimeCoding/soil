package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/supervisor"
	"os"
)

func New(o *option.Option) *Daemon {
	entries, e := procfile.Parse(o.ProcfilePath)
	errors.PanicOnError(e)
	v := environment.New(os.Environ())

	if f := v.Load(o.EnvrcPath); f != nil {
		errors.Printf("warning: %s\n", f)
	}

	s := supervisor.New(entries, v, o.ProcfilePath, o.EnvrcPath)

	return &Daemon{supervisor: s, service: service.New(s)}
}
