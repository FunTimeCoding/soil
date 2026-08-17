package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/pool"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Pool(v *pool.Pool) *server.Pool {
	return &server.Pool{
		Identifier:   v.Identifier,
		Interface:    v.Interface,
		StartAddress: v.StartAddress,
		EndAddress:   v.EndAddress,
		LeaseTime:    v.LeaseTime,
		Domain:       v.Domain,
		Description:  v.Description,
	}
}
