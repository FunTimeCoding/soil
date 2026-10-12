package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestProjectsWithoutASearchListsEveryProject(t *testing.T) {
	c := mock_client.New()
	c.SeedProject(1, "group", "app")
	c.SeedProject(2, "group", "web")
	projects, e := service_tester.Service(c).Projects("")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, projects)
}

func TestProjectsWithASearchNarrows(t *testing.T) {
	c := mock_client.New()
	c.SeedProject(1, "group", "app")
	c.SeedProject(2, "group", "web")
	projects, e := service_tester.Service(c).Projects("web")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, projects)
	assert.String(t, "group/web", projects[0].CombinedName())
}
