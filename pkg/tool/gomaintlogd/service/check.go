package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
)

func check(e *entry.Entry) error {
	if e.Action == "" {
		return validation.New("action is required")
	}

	if e.User == "" {
		return validation.New("user is required")
	}

	return nil
}
