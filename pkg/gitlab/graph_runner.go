package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/response"
)

func (c *Client) GraphRunner(identifier int64) (*response.Runner, error) {
	result := response.New()
	e := c.Query(
		fmt.Sprintf(
			"query {runner(id: \"gid://gitlab/Ci::Runner/%d\") { id description status runnerType managers { nodes { systemId ipAddress version revision } } } }",
			identifier,
		),
		&result,
	)

	if e != nil {
		return nil, e
	}

	return result, nil
}
