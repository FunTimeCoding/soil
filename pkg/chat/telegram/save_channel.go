package telegram

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"
	"go.etcd.io/bbolt"
)

func (c *Client) saveChannel(
	identifier int64,
	name string,
) {
	if c.database == nil {
		return
	}

	for _, a := range c.channels {
		if a.Identifier == identifier {
			return
		}
	}

	c.database.MustUpdate(
		func(t *bbolt.Tx) error {
			h := channel.New()
			h.Name = name
			h.Identifier = identifier
			fmt.Printf("New channel: %+v\n", h)
			c.database.MustPutBytes(
				c.database.Bucket(t, constant.ChannelBucket),
				h.Name,
				h.Encode(),
			)
			c.channels = append(c.channels, h)

			return nil
		},
	)
}
