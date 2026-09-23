package authority_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
)

func NewLeaf(
	cluster *authority.Authority,
	common string,
	host []string,
) *material.Material {
	n := distinguished_name.New()
	n.CommonName = common
	r := issue_request.New()
	r.Kind = constant.KindServer
	r.Name = n
	r.Host = host
	r.ValidDay = constant.LeafValidityDay

	return cluster.Issue(r)
}
