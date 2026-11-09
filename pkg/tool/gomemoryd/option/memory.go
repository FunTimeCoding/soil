package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Memory struct {
	option.Web
	LitePath      string
	HiddenTag     string
	ReferenceRoot string
	Clock         func() time.Time
	Query         func() face.Client
}
