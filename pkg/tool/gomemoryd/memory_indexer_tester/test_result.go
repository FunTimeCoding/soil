package memory_indexer_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func TestResult(
	path string,
	body string,
	score float32,
) *client.SearchResult {
	return &client.SearchResult{
		Path:        path,
		Body:        &body,
		Score:       score,
		Collection:  "test",
		Title:       path,
		Hash:        "test",
		Source:      "test",
		VirtualPath: fmt.Sprintf("qmd://test/%s", path),
		FilePath:    fmt.Sprintf("test/%s", path),
	}
}
