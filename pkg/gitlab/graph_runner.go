package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/response"
)

func (c *Client) GraphRunner(identifier int64) *response.Runner {
	result := response.New()
	c.MustQuery(
		fmt.Sprintf(
			"query {runner(id: \"gid://gitlab/Ci::Runner/%d\") { id description status runnerType managers { nodes { systemId ipAddress version revision } } } }",
			identifier,
		),
		&result,
	)

	return result
}
