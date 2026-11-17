package message

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"unicode/utf8"
)

func (s *Service) watchLimit(text string) {
	characters := utf8.RuneCountInString(text)

	if !s.metric.ObserveDelivery(characters) {
		return
	}

	s.logger.Structured(
		constant.DeliveryOverLimit,
		constant.DeliveryCharacters,
		characters,
		constant.DeliveryLimitKey,
		constant.DeliveryCharacterLimit,
	)
}
