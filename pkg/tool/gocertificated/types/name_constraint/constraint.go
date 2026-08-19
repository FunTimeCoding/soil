package name_constraint

import "net"

type Constraint struct {
	PermittedDomain  []string
	PermittedAddress []*net.IPNet
}
