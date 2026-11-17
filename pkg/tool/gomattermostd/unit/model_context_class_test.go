package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/mattermost_server"
	"path/filepath"
	"testing"
)

func TestMissingPostIsAnsweredQuietlyOverModelContext(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	c.MustCallToolError(
		constant.GetThreadReplies,
		map[string]any{"post_id": "missing"},
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDownloadingIntoMissingDirectoryIsAnsweredQuietly(t *testing.T) {
	s := base.New(t, mattermost_server.File, base.WithConnector())
	c := s.Tools()
	directory := filepath.Join(t.TempDir(), "absent")
	assert.StringContains(
		t,
		fmt.Sprintf("directory does not exist: %s", directory),
		c.MustCallToolError(
			constant.DownloadFile,
			map[string]any{
				"file_id": "alfa",
				"path":    filepath.Join(directory, "alfa.txt"),
			},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUploadingMissingFileIsAnsweredQuietly(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	path := filepath.Join(t.TempDir(), "absent.txt")
	assert.StringContains(
		t,
		"no such file",
		c.MustCallToolError(
			constant.UploadFile,
			map[string]any{"path": path, "channel_id": "bravo"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}
