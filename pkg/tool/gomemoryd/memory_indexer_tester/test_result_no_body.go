package memory_indexer_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func TestResultNoBody(path string) *client.SearchResult {
	return &client.SearchResult{
		Path:        path,
		Score:       0,
		Collection:  "test",
		Title:       path,
		Hash:        "test",
		Source:      "test",
		VirtualPath: fmt.Sprintf("qmd://test/%s", path),
		FilePath:    fmt.Sprintf("test/%s", path),
	}
}
