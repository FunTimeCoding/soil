package ladybird

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/request"
)

func (c *Client) CallFunction(
	context string,
	declaration string,
	arguments ...any,
) (json.RawMessage, error) {
	return scriptResult(
		c.send(
			constant.CallMethod,
			request.NewCall(context, declaration, arguments),
		),
	)
}
