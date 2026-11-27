package ladybird

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/request"
)

func (c *Client) Evaluate(
	context string,
	expression string,
) (json.RawMessage, error) {
	return scriptResult(
		c.send(
			constant.LadybirdEvaluateMethod,
			request.NewEvaluate(context, expression),
		),
	)
}
