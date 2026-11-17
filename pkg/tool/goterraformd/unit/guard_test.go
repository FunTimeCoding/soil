package unit

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, mock_client.New()).Served()
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/trigger")
	v.VerifyModelContext(t)
}
