package internet_address

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"log"
	"slices"
)

func validateObjectType(objectType string) {
	if !slices.Contains(constant.InternetAddressTargets, objectType) {
		log.Panicf("unexpected object type: %s", objectType)
	}
}
