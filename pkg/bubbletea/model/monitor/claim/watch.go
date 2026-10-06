package claim

import (
	"context"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"time"
)

func Watch(
	c *client.Client,
	updates chan Message,
) {
	for {
		e := c.Stream(
			context.Background(),
			func(v []generated.Claim) {
				updates <- Message{Claims: v}
			},
		)
		updates <- Message{Error: e}
		time.Sleep(constant.ReconnectDelay)
	}
}
