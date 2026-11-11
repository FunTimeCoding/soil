package document

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
	"time"
)

func completionMetadata(
	metadata map[string][]string,
	c *completion.Completion,
) map[string][]string {
	metadata["created_at"] = []string{c.CreatedAt.UTC().Format(time.RFC3339)}
	metadata[constant.CompletedAtMetadata] = []string{
		c.CompletedAt.UTC().Format(time.RFC3339),
	}

	return metadata
}
