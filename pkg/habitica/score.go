package habitica

import (
	"github.com/funtimecoding/soil/pkg/habitica/score"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) Score(
	taskIdentifier string,
	direction string,
) (*score.Score, error) {
	if direction == "" {
		direction = "up"
	}

	var result *score.Score

	return result, c.post(
		join.Empty("/tasks/", taskIdentifier, "/score/", direction),
		nil,
		&result,
	)
}
