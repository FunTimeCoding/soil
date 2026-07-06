package habitica

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/habitica/score"
)

func (c *Client) MustScore(
	taskIdentifier string,
	direction string,
) *score.Score {
	result, e := c.Score(taskIdentifier, direction)
	errors.PanicOnError(e)

	return result
}
