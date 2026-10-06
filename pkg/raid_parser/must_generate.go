package raid_parser

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustGenerate(
	files []string,
	date *time.Time,
) string {
	result, e := c.Generate(files, date)
	errors.PanicOnError(e)

	return result
}
