package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/new_session"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) createSession() *new_session.Response {
	body := notation.Encode(
		map[string]any{
			"capabilities": map[string]any{
				"alwaysMatch": map[string]any{
					constant.LadybirdWebSocketCapability: true,
				},
			},
		},
		false,
	)
	r := c.driverSend(
		web.NewPost(c.driverLocator(constant.LadybirdSessionRoute), body),
	)
	result := new_session.NewResponse()
	errors.PanicOnError(notation.DecodeBytes(web.ReadBytes(r), result))

	return result
}
