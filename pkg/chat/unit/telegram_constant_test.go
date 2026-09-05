package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "TELEGRAM_TOKEN", constant.TelegramTokenEnvironment)
}
