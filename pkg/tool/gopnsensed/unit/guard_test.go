package unit

import (
	"github.com/funtimecoding/soil/pkg/opnsense/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, mock_client.New(t)).Served()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
