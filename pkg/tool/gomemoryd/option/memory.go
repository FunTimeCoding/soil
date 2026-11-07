package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Memory struct {
	option.Web
	LitePath      string
	HiddenTag     string
	ReferenceRoot string
	Query         func() face.Client
}
