package physical_address

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"log"
	"slices"
)

func validateObjectType(objectType string) {
	if !slices.Contains(constant.PhysicalObjectTypes, objectType) {
		log.Panicf("unexpected object type: %s", objectType)
	}
}
