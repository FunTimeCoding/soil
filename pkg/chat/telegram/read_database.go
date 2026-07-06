package telegram

import (
	"github.com/funtimecoding/soil/pkg/bolt"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/user"
	"go.etcd.io/bbolt"
)

func (c *Client) readDatabase() {
	if c.database == nil {
		return
	}

	c.database.MustView(
		func(t *bbolt.Tx) error {
			bolt.MustFor(
				c.database.Bucket(t, constant.ChannelBucket),
				func(
					k string,
					v []byte,
				) {
					a := channel.New()
					a.Decode(v)
					c.channels = append(c.channels, a)
				},
			)
			bolt.MustFor(
				c.database.Bucket(t, constant.UserBucket),
				func(
					k string,
					v []byte,
				) {
					s := user.New()
					s.Decode(v)
					c.users = append(c.users, s)
				},
			)

			return nil
		},
	)
}
