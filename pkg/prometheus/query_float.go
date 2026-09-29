package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"time"
)

func (c *Client) QueryFloat(
	q string,
	t time.Time,
) (float64, error) {
	result, e := c.QueryFloats(q, t)

	if e != nil {
		return 0, e
	}

	if len(result) > 1 {
		return 0, validation.New(
			"query %s returned %d series, expected one",
			q,
			len(result),
		)
	}

	for _, v := range result {
		return v, nil
	}

	return 0, nil
}
