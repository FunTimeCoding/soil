package logger

import "github.com/funtimecoding/soil/pkg/web/request_context"

func (l *Logger) Request(c *request_context.Context) {
	c.LogWebhook(l.context, l.structured)
}
