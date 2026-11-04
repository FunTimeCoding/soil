package message

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"unicode/utf8"
)

func (s *Service) watchBudget(text string) {
	characters := utf8.RuneCountInString(text)

	if !s.metric.ObserveDelivery(characters) {
		return
	}

	s.logger.Structured(
		constant.DeliveryOverBudget,
		constant.DeliveryCharacters,
		characters,
		constant.DeliveryBudgetKey,
		constant.DeliveryBudget,
	)
}
