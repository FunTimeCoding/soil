package unit

import "github.com/funtimecoding/soil/pkg/gitlab/mock_client"

func seedForge(
	c *mock_client.Client,
	project int64,
	name string,
) {
	c.SeedProject(project, "group", name)
	c.SeedBranch("main", "aaa")
	c.SeedPipeline(10, "main", "aaa", "success")
}
