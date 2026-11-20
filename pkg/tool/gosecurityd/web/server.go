package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/overview"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	vulnerability *vulnerability.Service
	debian        *debian.Service
	flag          *flag.Service
	overview      *overview.Service
	system        *system.Service
	reporter      face.Reporter
	view          *view.View
}
