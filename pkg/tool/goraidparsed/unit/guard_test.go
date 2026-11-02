package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/fake_parser"
	"testing"
)

func TestGuard(t *testing.T) {
	v := mounted(t, fake_parser.New(0, nil))
	v.VerifyBase(t)
}
