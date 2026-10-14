package unit

import (
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := startServer(t, service_tester.Single(mock_client.New()))
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyOpen(t, constant.BoardPath)
	v.VerifyOpen(t, "/event")
	v.VerifyModelContext(t)
}
