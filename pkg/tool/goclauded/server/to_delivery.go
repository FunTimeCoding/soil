package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func toDelivery(d *server.Delivery) constant.Delivery {
	if d == nil {
		return constant.DeliveryQueued
	}

	return constant.Delivery(*d)
}
