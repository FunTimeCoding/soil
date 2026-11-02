package service_tester

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"path/filepath"
	"testing"
)

func NewOption(
	t *testing.T,
	c chat.MattermostClient,
	k goclauded.Client,
	configure func(*option.Mattermost),
) *option.Mattermost {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "mattermost.db")
	result.Mattermost = func() chat.MattermostClient { return c }
	result.Connector = func() goclauded.Client { return k }

	if configure != nil {
		configure(result)
	}

	return result
}
