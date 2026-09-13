package mechanic

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	n face.EventNotifier,
	extended string,
	serverSide string,
) *Server {
	return &Server{
		notifier: n,
		view: view.New(
			layout.New(identity.Example()).
				WithExtended(extended).
				WithServerSide(serverSide).
				WithLiveEndpoint(constant.LivePath).
				WithItems(
					navigation_item.New(constant.RootPath, constant.SwapTitle),
					navigation_item.New(
						constant.TriggerPath,
						constant.TriggerTitle,
					),
					navigation_item.New(
						constant.NotifyPath,
						constant.NotifyTitle,
					),
					navigation_item.New(
						constant.StreamPath,
						constant.StreamTitle,
					),
					navigation_item.New(
						constant.ExtraPath,
						constant.ExtraTitle,
					),
				),
		),
	}
}
