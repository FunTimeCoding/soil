package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"unicode/utf8"
)

func entryRefusal(
	kind string,
	body string,
) error {
	characters := utf8.RuneCountInString(body)

	if characters <= constant.DeliveryEntryLimit {
		return nil
	}

	return refusal.New(
		fmt.Sprintf(
			constant.DeliveryRefused,
			kind,
			characters,
			constant.DeliveryEntryLimit,
		),
	)
}
