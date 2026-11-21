package remediation

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
)

type Service struct {
	system *system.Service
	flag   *flag.Service
}
