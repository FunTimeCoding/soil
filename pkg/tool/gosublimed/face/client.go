package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

type Client interface {
	CloseView(
		x context.Context,
		identifier int,
	) error
	CreateView(
		x context.Context,
		b client.CreateViewJSONRequestBody,
	) (*client.View, error)
	EditView(
		x context.Context,
		identifier int,
		b client.EditViewJSONRequestBody,
	) (*client.View, error)
	GetView(
		x context.Context,
		identifier int,
	) (*client.View, error)
	GetViews(x context.Context) ([]client.View, error)
	OpenFile(
		x context.Context,
		b client.OpenFileJSONRequestBody,
	) (*client.View, error)
	SaveView(
		x context.Context,
		identifier int,
		b client.SaveViewJSONRequestBody,
	) error
	Reachable(x context.Context) error
	Close()
}
