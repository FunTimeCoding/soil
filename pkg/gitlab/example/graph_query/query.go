package graph_query

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

// Reference: https://docs.gitlab.com/api/graphql/
func Query() {
	c := gitlab.NewEnvironment()
	runner := &RunnerResult{}
	c.MustQuery(
		`query {
  runner(id: "gid://gitlab/Ci::Runner/1") {
    id description status runnerType
    managers { nodes { systemId ipAddress version revision } }
  }
}`,
		&runner,
	)
	console.Format("Response: %+v\n", runner)
	console.Format("Runner: %+v\n", c.MustGraphRunner(1))
}
