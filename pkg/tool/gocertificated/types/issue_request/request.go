package issue_request

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/name_constraint"
	"time"
)

type Request struct {
	Kind       constant.CertificateKind
	Name       *distinguished_name.Name
	Constraint *name_constraint.Constraint
	Host       []string
	Start      time.Time
	ValidYear  int
	ValidDay   int
}
