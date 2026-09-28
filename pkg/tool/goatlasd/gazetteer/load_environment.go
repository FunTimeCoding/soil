package gazetteer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/netbox"
)

func LoadEnvironment(q context.Context) (*Gazetteer, error) {
	return Load(q, netbox.NewEnvironment())
}
