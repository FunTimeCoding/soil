package outpost

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"strings"
)

func resolveHost(
	s *gazetteer.Gazetteer,
	host *client.Host,
) (*place.Place, bool) {
	for _, name := range []string{
		host.Hostname,
		strings.SplitN(host.Hostname, constant.Dot, 2)[0],
	} {
		if v, okay := s.Resolve(name); okay {
			return v, true
		}
	}

	if host.HardwareAddresses == nil {
		return nil, false
	}

	for _, address := range *host.HardwareAddresses {
		if v, okay := s.ResolveHardware(address); okay {
			return v, true
		}
	}

	return nil, false
}
