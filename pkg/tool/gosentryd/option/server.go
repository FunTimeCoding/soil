package option

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Server struct {
	option.Web
	Organization string
	Host         string
	Source       func() face.SentryClient
}
