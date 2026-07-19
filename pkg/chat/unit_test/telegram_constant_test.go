package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "TELEGRAM_TOKEN", constant.TokenEnvironment)
}
