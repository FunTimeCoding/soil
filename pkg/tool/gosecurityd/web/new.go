package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/overview"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *vulnerability.Service,
	d *debian.Service,
	f *flag.Service,
	o *overview.Service,
	y *system.Service,
	r face.Reporter,
) *Server {
	return &Server{
		vulnerability: v,
		debian:        d,
		flag:          f,
		overview:      o,
		system:        y,
		reporter:      r,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeSentinel, favicon).
				WithItems(
					navigation_item.New(
						constant.SystemsPath,
						constant.SystemsTitle,
					),
					navigation_item.New(
						constant.LookupPath,
						constant.LookupTitle,
					),
				),
		),
	}
}
