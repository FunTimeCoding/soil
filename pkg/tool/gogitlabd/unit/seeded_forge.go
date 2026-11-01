package unit

import (
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"testing"
)

func seededForge(
	t *testing.T,
	project int64,
	name string,
) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)
	result.SeedProject(project, "group", name)
	result.SeedBranch("main", "aaa")
	result.SeedPipeline(10, "main", "aaa", "success")

	return result
}
