package console_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"testing"
)

type Tester struct {
	t      *testing.T
	client face.Client
}
