package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/model_context_tester"
	"path/filepath"
	"testing"
)

func TestAMissingPostIsAnsweredQuietlyOverModelContext(t *testing.T) {
	r := model_context_tester.New(t, threadUpstream)
	r.Client.MustCallToolError(
		constant.GetThreadReplies,
		map[string]any{"post_id": "missing"},
	)
	assert.Count(t, 0, r.Server.Reporter.Events())
}

func TestDownloadingIntoAMissingDirectoryIsAnsweredQuietly(t *testing.T) {
	r := model_context_tester.New(t, fileUpstream)
	directory := filepath.Join(t.TempDir(), "absent")
	assert.StringContains(
		t,
		fmt.Sprintf("directory does not exist: %s", directory),
		r.Client.MustCallToolError(
			constant.DownloadFile,
			map[string]any{
				"file_id": "alfa",
				"path":    filepath.Join(directory, "alfa.txt"),
			},
		),
	)
	assert.Count(t, 0, r.Server.Reporter.Events())
}

func TestUploadingAMissingFileIsAnsweredQuietly(t *testing.T) {
	r := model_context_tester.New(t, threadUpstream)
	path := filepath.Join(t.TempDir(), "absent.txt")
	assert.StringContains(
		t,
		"no such file",
		r.Client.MustCallToolError(
			constant.UploadFile,
			map[string]any{"path": path, "channel_id": "bravo"},
		),
	)
	assert.Count(t, 0, r.Server.Reporter.Events())
}
