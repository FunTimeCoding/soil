package remediation

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
)

func New(
	y *system.Service,
	f *flag.Service,
) *Service {
	return &Service{system: y, flag: f}
}
