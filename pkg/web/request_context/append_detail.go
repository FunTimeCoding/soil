package request_context

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/text/multi_line"
)

func AppendDetail(
	l *multi_line.MultiLine,
	c *Context,
) {
	l.Format("Raw query: %s", c.Request().URL.RawQuery)

	for k, v := range c.Request().Header {
		l.Format("Header: %s=%s", k, join.Comma(v))
	}

	l.Format("Body: %s", c.Body())
}
