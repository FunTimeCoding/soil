package subscription

import goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s Subscription) DeliveryMode() goclaudedConstant.Delivery {
	if s.Delivery == "" {
		return goclaudedConstant.DeliveryImmediate
	}

	return goclaudedConstant.Delivery(s.Delivery)
}
