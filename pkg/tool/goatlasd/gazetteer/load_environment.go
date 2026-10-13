package gazetteer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/netbox"
	"net/http"
)

func LoadEnvironment(
	q context.Context,
	h *http.Client,
) (*Gazetteer, error) {
	return Load(q, netbox.NewEnvironment(h))
}
