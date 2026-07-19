package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/telegram/message"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestNewSlice(t *testing.T) {
	assert.Count(t, 0, message.NewSlice([]*tgbotapi.Message{}))
}
