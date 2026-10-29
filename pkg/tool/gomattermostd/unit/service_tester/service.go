package service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"testing"
)

func Service(
	t *testing.T,
	c face.Client,
	k *connector.Client,
	configure func(*option.Mattermost),
) *service.Service {
	t.Helper()

	return Reported(t, c, k, configure, memory.New())
}
