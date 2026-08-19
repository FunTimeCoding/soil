package authority

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/policy"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
)

func NewRoot(
	name *distinguished_name.Name,
	year int,
) *Authority {
	policy.New(constant.KindRoot).Validate(name)
	r := issue_request.New()
	r.Kind = constant.KindRoot
	r.Name = name
	r.ValidYear = year
	t := newTemplate(r)
	k := newKey()

	return New(material.New(newCertificate(t, t, k.Public(), k), k))
}
