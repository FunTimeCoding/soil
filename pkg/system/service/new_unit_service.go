package service

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
)

func NewUnitService(
	unit string,
	state string,
	path string,
	name string,
	policies map[string]*Policy,
	manual map[string]bool,
) *Service {
	if name == "" {
		return NewService(
			unit,
			state,
			constant.ServiceOriginLocal,
			path,
			"",
			"",
			true,
		)
	}

	policy := policies[name]

	if policy == nil {
		policy = &Policy{}
	}

	origin := constant.ServiceOriginVendor

	if strings.HasPrefix(policy.Origin, constant.OriginMirrorPrefix) {
		origin = constant.ServiceOriginSystem
	}

	return NewService(
		unit,
		state,
		origin,
		policy.Origin,
		name,
		policy.Version,
		manual[name],
	)
}
