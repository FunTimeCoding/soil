package logger

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/constant"
	"log/slog"
	"time"
)

func (l *Logger) WebSocket(v *message.Message) {
	if v.Event == nil {
		return
	}

	t, e := time.Parse(time.RFC3339Nano, v.Event.Time)
	errors.PanicOnError(e)
	result := []slog.Attr{
		slog.String(constant.MessagingSystem, "homeassistant"),
		slog.String(constant.OperationName, "receive"),
		slog.String(
			constant.MessageIdentifier,
			fmt.Sprintf("%d", v.Identifier),
		),
		slog.String(constant.MessageType, v.Event.Type),
		slog.String(
			constant.ConversationIdentifier,
			v.Event.Context.Identifier,
		),
		slog.Time(constant.EnvelopeTime, t),
		slog.Int(constant.BodySize, len(v.Event.Raw)),
	}

	if v.Event.Origin != "" {
		result = append(
			result,
			slog.String(constant.HomeAssistantOrigin, v.Event.Origin),
		)
	}

	l.structured.LogAttrs(
		l.context,
		slog.LevelInfo,
		"websocket_event_receive",
		result...,
	)
}
