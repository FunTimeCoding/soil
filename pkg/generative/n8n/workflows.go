package n8n

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/n8n/response"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) Workflows() []*response.Workflow {
	var r *response.Workflows
	notation.MustDecode(c.Get(constant.N8nWorkflows), &r, false)

	return r.Payload
}
