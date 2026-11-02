package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (o *Tester) UnreachableInventory() {
	o.Option.Inventory = func() face.Client {
		return client.New(
			locator.New("localhost").Port(1).Insecure(),
			constant.FixtureToken,
		)
	}
}
