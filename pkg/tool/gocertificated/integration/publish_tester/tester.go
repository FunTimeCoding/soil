package publish_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/integration/base"
	"testing"
)

type Tester struct {
	Server *base.Server
	t      *testing.T
}
