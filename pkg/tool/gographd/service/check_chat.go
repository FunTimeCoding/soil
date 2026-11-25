package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"strings"
)

func checkChat(chat string) error {
	if !constant.ChatPattern.MatchString(chat) ||
		strings.Contains(chat, constant.ParentStep) {
		return validation.New(
			"chat identifier has an unexpected shape: %s",
			chat,
		)
	}

	return nil
}
