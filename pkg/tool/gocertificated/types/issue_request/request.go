package issue_request

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/name_constraint"
)

type Request struct {
	Kind       constant.CertificateKind
	Name       *distinguished_name.Name
	Constraint *name_constraint.Constraint
	Host       []string
	ValidYear  int
	ValidDay   int
}
