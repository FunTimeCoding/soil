package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/service"
)

func New(o *option.Alpine) *Daemon {
	return &Daemon{
		service: service.New(
			o.PackageRoot,
			o.KeyDirectory,
			o.SignatureKey,
			o.Indexer(),
		),
		packageRoot: o.PackageRoot,
	}
}
