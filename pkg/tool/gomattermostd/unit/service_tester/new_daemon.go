package service_tester

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"testing"
)

func NewDaemon(
	t *testing.T,
	c chat.MattermostClient,
	k goclauded.Client,
	configure func(*option.Mattermost),
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(
		NewOption(t, c, k, configure),
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(result.Close)

	return result
}
