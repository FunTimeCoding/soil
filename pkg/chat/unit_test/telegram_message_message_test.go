package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/telegram/message"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestMessage(t *testing.T) {
	assert.NotNil(
		t,
		message.New(
			&tgbotapi.Message{From: &tgbotapi.User{UserName: upper.Alfa}},
		),
	)
}
