package service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"testing"
)

func Monitor(
	t *testing.T,
	m face.MattermostClient,
) *service.Service {
	t.Helper()

	return Service(
		t,
		m,
		nil,
		func(o *option.Mattermost) { o.Monitor.Enabled = true },
	)
}
