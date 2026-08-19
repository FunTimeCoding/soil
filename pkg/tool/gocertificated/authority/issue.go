package authority

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
)

func (a *Authority) Issue(r *issue_request.Request) *material.Material {
	k := newKey()

	return material.New(a.Sign(r, k.Public()), k)
}
