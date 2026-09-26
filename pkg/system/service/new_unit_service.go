package service

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
)

func newUnitService(
	unit string,
	unitState string,
	path string,
	name string,
	policies map[string]*Policy,
	manual map[string]bool,
) *Service {
	if name == "" {
		return NewService(
			unit,
			unitState,
			constant.ServiceOriginLocal,
			path,
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
		unitState,
		origin,
		policy.Origin,
		policy.Version,
		manual[name],
	)
}
