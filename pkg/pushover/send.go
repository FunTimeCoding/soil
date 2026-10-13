package pushover

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/pushover/constant"
	"github.com/funtimecoding/soil/pkg/pushover/notification"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func (c *Client) Send(s string) *http.Response {
	result, e := c.pool.Client().Post(
		locator.New(constant.Host).Base(constant.Base).Path(constant.Message).String(),
		webConstant.Object,
		bytes.NewBuffer(
			[]byte(
				notation.Encode(notification.New(c.user, c.token, s), false),
			),
		),
	)
	errors.PanicOnError(e)

	return result
}
