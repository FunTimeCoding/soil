package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	"testing"
)

func Service(t *testing.T) *service.Service {
	t.Helper()

	return New(t).Service
}
