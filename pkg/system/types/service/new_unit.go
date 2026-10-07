package service

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/types/service_policy"
	"strings"
)

func NewUnit(
	unit string,
	state string,
	path string,
	name string,
	policies map[string]*service_policy.Policy,
	manual map[string]bool,
) *Service {
	if name == "" {
		return New(unit, state, constant.ServiceOriginLocal, path, "", "", true)
	}

	p := policies[name]

	if p == nil {
		p = service_policy.New()
	}

	origin := constant.ServiceOriginVendor

	if strings.HasPrefix(p.Origin, constant.OriginMirrorPrefix) {
		origin = constant.ServiceOriginSystem
	}

	return New(unit, state, origin, p.Origin, name, p.Version, manual[name])
}
