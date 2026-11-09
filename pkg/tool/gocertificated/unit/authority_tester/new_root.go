package authority_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"time"
)

func NewRoot() *authority.Authority {
	return authority.NewRoot(
		NewRootName(),
		constant.RootValidityYear,
		time.Now(),
	)
}
