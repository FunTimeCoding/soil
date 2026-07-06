package telegram

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/user"
	"go.etcd.io/bbolt"
)

func (c *Client) saveUser(
	identifier int64,
	name string,
) {
	if c.database == nil {
		return
	}

	for _, s := range c.users {
		if s.Identifier == identifier {
			return
		}
	}

	c.database.MustUpdate(
		func(t *bbolt.Tx) error {
			u := user.New()
			u.Name = name
			u.Identifier = identifier
			fmt.Printf("New user: %+v\n", u)
			c.database.MustPutBytes(
				c.database.Bucket(t, constant.UserBucket),
				u.Name,
				u.Encode(),
			)
			c.users = append(c.users, u)

			return nil
		},
	)
}
