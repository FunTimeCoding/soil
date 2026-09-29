package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"time"
)

func (c *Client) QueryInteger(
	q string,
	t time.Time,
) (int, error) {
	result, e := c.QueryIntegers(q, t)

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
