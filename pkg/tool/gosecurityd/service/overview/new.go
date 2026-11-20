package overview

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
)

func New(
	f *flag.Service,
	d *debian.Service,
) *Service {
	return &Service{flag: f, debian: d}
}
