package network

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
	"log"
	"slices"
)

func validateType(t netbox.InterfaceTypeValue) {
	if !slices.Contains(constant.InterfaceTypes, t) {
		log.Panicf("unexpected interface type: %s", t)
	}
}
