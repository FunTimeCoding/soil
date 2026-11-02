package service_tester

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"testing"
)

func Service(
	t *testing.T,
	c chat.MattermostClient,
	k goclauded.Client,
	configure func(*option.Mattermost),
) *service.Service {
	t.Helper()

	return Reported(t, c, k, configure, memory.New())
}
