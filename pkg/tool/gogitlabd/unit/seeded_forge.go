package unit

import "github.com/funtimecoding/soil/pkg/gitlab/mock_client"

func seededForge(
	project int64,
	name string,
) *mock_client.Client {
	result := mock_client.New()
	result.SeedProject(project, "group", name)
	result.SeedBranch("main", "aaa")
	result.SeedPipeline(10, "main", "aaa", "success")

	return result
}
