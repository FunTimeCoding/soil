package overview

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
)

type Service struct {
	flag   *flag.Service
	debian *debian.Service
}
