package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"strings"
)

func checkDescription(description string) error {
	if strings.ContainsAny(description, "\r\n") {
		return validation.New(
			"description must be one line; the body belongs in content",
		)
	}

	return nil
}
