package message

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
)

func delivery(
	d constant.Delivery,
	absent constant.Delivery,
) (constant.Delivery, error) {
	switch d {
	case "":
		return absent, nil
	case constant.DeliveryQueued, constant.DeliveryImmediate:
		return d, nil
	default:
		return "", refusal.New(constant.DeliveryUnknown)
	}
}
