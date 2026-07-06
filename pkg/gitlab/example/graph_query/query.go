package graph_query

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

/*
Working

query {
  users(admins: true) { nodes { id username name publicEmail webUrl } }
}

query {
  runner(id: "gid://gitlab/Ci::Runner/1") {
    id description status runnerType
    managers { nodes { systemId ipAddress version revision } }
  }
}

query { project(fullPath: "gw2/wvw-schedule") { id } }

query { project(fullPath: "gw2/wvw-schedule") { id } }
*/

/*
Not working

query { runners(first: 5) { nodes { id description } } }
*/

func Query() {
	// Reference:
	// https://docs.gitlab.com/api/graphql/
	// TODO: How to get all runners?
	// https://docs.gitlab.com/api/graphql/reference/#queryrunners
	// Explorer, also exists on self-hosted:
	// https://gitlab.com/-/graphql-explorer
	c := gitlab.NewEnvironment()
	runner := &RunnerResult{}
	// Not sure how to get all runners, but this works, and IDs can be loaded the other API way
	c.MustQuery(
		`query {
  runner(id: "gid://gitlab/Ci::Runner/1") {
    id description status runnerType
    managers { nodes { systemId ipAddress version revision } }
  }
}`,
		&runner,
	)
	fmt.Printf("Response: %+v\n", runner)
	fmt.Printf("Runner: %+v\n", c.GraphRunner(1))
}
