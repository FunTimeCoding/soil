package service

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"unicode/utf8"
)

func (s *Service) watchBudget(text string) {
	characters := utf8.RuneCountInString(text)

	if characters <= constant.DeliveryBudget {
		return
	}

	s.reporter.CaptureWithContext(
		errors.New(constant.DeliveryOverBudget),
		constant.DeliveryContextKey,
		map[string]any{constant.DeliveryCharacters: characters},
	)
}
