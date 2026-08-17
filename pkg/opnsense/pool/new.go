package pool

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.Pool) *Pool {
	return &Pool{
		Identifier:   v.Identifier,
		Interface:    v.Interface,
		StartAddress: v.StartAddress,
		EndAddress:   v.EndAddress,
		LeaseTime:    v.LeaseTime,
		Domain:       v.Domain,
		Description:  v.Description,
	}
}
