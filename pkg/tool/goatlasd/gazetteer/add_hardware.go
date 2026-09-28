package gazetteer

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"strings"
)

func (s *Gazetteer) addHardware(v *client.PhysicalAddressOwner) {
	if v.ObjectKind == nil || v.ObjectIdentifier == nil ||
		v.ObjectName == nil {
		return
	}

	s.hardware[strings.ToLower(v.Address)] = place.New(
		*v.ObjectKind,
		*v.ObjectIdentifier,
		*v.ObjectName,
	)
}
