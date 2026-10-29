package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/option"
)

func New(
	o *option.Directory,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := o.Source()

	return &Service{
		directory:     d,
		guard:         peer.New(constant.DirectoryPeer, o.Host, d, l, r),
		authorization: authorizationClient(o).WithReporter(r),
	}
}
