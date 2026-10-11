package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"time"
)

func messageEntry(m *message.Message) *server.MessageEntry {
	return &server.MessageEntry{
		Identifier: int(m.Identifier),
		From:       m.FromName,
		To:         m.ToName,
		Body:       m.Body,
		Timestamp:  m.CreatedAt.UTC().Format(time.RFC3339),
	}
}
