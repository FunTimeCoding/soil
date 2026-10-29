package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
)

func New(
	o *option.Log,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	a := o.Alertmanager()

	return &Service{
		database:     d,
		store:        store.New(d),
		alertmanager: a,
		guard:        peer.New(constant.AlertmanagerPeer, a.Address(), a, l, r),
		retention:    o.Retention,
	}
}
