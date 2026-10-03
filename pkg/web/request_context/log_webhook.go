package request_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"log/slog"
)

func (c *Context) LogWebhook(
	x context.Context,
	l *slog.Logger,
) {
	l.LogAttrs(
		x,
		slog.LevelInfo,
		constant.RequestStart,
		c.WebhookAttribute()...,
	)
}
