package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/supervisor"
	"os"
)

func New(o *option.Option) *Service {
	entries, e := procfile.Parse(o.ProcfilePath)
	errors.PanicOnError(e)
	v := environment.New(os.Environ())

	if f := v.Load(o.EnvrcPath); f != nil {
		errors.Printf("warning: %s\n", f)
	}

	return &Service{
		supervisor: supervisor.New(entries, v, o.ProcfilePath, o.EnvrcPath),
	}
}
