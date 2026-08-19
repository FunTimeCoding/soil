package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/name_constraint"
	"net"
)

func constraint(
	domain []string,
	address []string,
) (*name_constraint.Constraint, error) {
	if len(domain) == 0 && len(address) == 0 {
		return nil, nil
	}

	result := name_constraint.New()
	result.PermittedDomain = domain

	for _, a := range address {
		_, network, e := net.ParseCIDR(a)

		if e != nil {
			return nil, errors.Format("permitted address is not a network", a)
		}

		result.PermittedAddress = append(result.PermittedAddress, network)
	}

	return result, nil
}
