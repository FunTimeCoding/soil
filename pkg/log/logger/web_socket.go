package logger

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/errors"
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
		slog.String(MessagingSystem, "homeassistant"),
		slog.String(OperationName, "receive"),
		slog.String(
			MessageIdentifier,
			fmt.Sprintf("%d", v.Identifier),
		),
		slog.String(MessageType, v.Event.Type),
		slog.String(
			ConversationIdentifier,
			v.Event.Context.Identifier,
		),
		slog.Time(EnvelopeTime, t),
		slog.Int(BodySize, len(v.Event.Raw)),
	}

	if v.Event.Origin != "" {
		result = append(
			result,
			slog.String(HomeAssistantOrigin, v.Event.Origin),
		)
	}

	l.structured.LogAttrs(
		l.context,
		slog.LevelInfo,
		"websocket_event_receive",
		result...,
	)
}
