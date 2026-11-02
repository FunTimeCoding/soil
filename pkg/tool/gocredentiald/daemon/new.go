package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/option"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"time"
)

func New(
	o *option.Credential,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	return &Daemon{
		service: service.New(
			keepass.MustNew(o.Database, o.Password),
			o.Database,
			o.Password,
			o.RevealedField,
			time.Now,
			l,
			r,
		),
	}
}
