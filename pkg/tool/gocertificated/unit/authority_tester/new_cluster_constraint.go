package authority_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/name_constraint"
	"net"
)

func NewClusterConstraint() *name_constraint.Constraint {
	_, permitted, e := net.ParseCIDR(constant.FixtureAddress)

	if e != nil {
		panic(e)
	}

	c := name_constraint.New()
	c.PermittedDomain = []string{
		constant.FixtureDomain,
		constant.FixtureInternalDomain,
		constant.FixtureLocalDomain,
	}
	c.PermittedAddress = []*net.IPNet{permitted}

	return c
}
