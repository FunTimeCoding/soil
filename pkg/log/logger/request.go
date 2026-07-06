package logger

import "github.com/funtimecoding/soil/pkg/web/request_context"

func (l *Logger) Request(c *request_context.Context) {
	c.LogStart(l.context, l.structured)
}
