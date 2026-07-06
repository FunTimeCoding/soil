package message

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestNewSlice(t *testing.T) {
	assert.Count(t, 0, NewSlice([]*tgbotapi.Message{}))
}
