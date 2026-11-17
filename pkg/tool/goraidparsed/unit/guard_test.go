package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/mock_parser"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, mock_parser.New(0, nil)).Served()
	v.VerifyBase(t)
}
