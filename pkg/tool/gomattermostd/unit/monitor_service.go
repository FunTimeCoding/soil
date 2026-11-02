package unit

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"testing"
)

func monitorService(
	t *testing.T,
	m *mock_client.Client,
) *service.Service {
	t.Helper()

	return service_tester.Service(
		t,
		m,
		nil,
		func(o *option.Mattermost) { o.Monitor.Enabled = true },
	)
}
