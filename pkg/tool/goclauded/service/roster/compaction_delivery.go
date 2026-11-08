package roster

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func compactionDelivery(p constant.CompactionPhase) (constant.Delivery, error) {
	switch p {
	case constant.CompactionStarted:
		return constant.DeliveryQueued, nil
	case constant.CompactionFinished:
		return constant.DeliveryImmediate, nil
	}

	return "", validation.New(constant.CompactionUnknown, p)
}
