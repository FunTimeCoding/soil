package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/face"
	"testing"
)

type Tester struct {
	t        *testing.T
	Service  face.Ansible
	Reporter *memory.Memory
	capture  string
}
