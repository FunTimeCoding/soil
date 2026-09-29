package prometheus

import "time"

func (c *Client) IntegerExists(
	q string,
	t time.Time,
) (bool, error) {
	result, e := c.QueryIntegers(q, t)

	if e != nil {
		return false, e
	}

	return len(result) > 0, nil
}
