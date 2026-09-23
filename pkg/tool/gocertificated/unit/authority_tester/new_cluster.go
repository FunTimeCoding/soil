package authority_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
)

func NewCluster(root *authority.Authority) *authority.Authority {
	n := distinguished_name.New()
	n.CommonName = constant.FixtureIssuingCommonName
	r := issue_request.New()
	r.Kind = constant.KindIntermediate
	r.Name = n
	r.Constraint = NewClusterConstraint()
	r.ValidYear = constant.IntermediateValidityYear

	return authority.New(root.Issue(r))
}
