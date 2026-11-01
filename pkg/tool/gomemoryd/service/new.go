package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func New(
	o *option.Memory,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := connection.New(l, o.LitePath)
	q := o.Query()

	return &Service{
		reporter:      r,
		database:      d,
		store:         store.New(d),
		query:         q,
		guard:         peer.New(constant.Identity.Name(), q.Address(), q, l, r),
		hiddenTag:     o.HiddenTag,
		referenceRoot: o.ReferenceRoot,
	}
}
