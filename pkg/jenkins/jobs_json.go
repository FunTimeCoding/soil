package jenkins

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) JobsJSON() {
	exec := gojenkins.Executor{
		Raw:     new(gojenkins.ExecutorResponse),
		Jenkins: c.client,
	}
	_, e := c.client.Requester.GetJSON(
		c.context,
		"/",
		exec.Raw,
		nil,
	)
	errors.PanicOnError(e)
	jobs := make([]*gojenkins.Job, len(exec.Raw.Jobs))
	jobStrings := make([]string, len(exec.Raw.Jobs))

	for i, job := range exec.Raw.Jobs {
		res := new(gojenkins.JobResponse)
		r2, f := c.client.Requester.GetJSON(
			c.context,
			fmt.Sprintf("/job/%s", job.Name),
			res,
			nil,
		)
		errors.PanicOnError(f)

		if false {
			// TODO: Body is already read, would have to craft entire request
			jobStrings[i] = web.ReadString(r2)
		}

		if false {
			ji, g := c.client.GetJob(c.context, job.Name)
			errors.PanicOnError(g)
			jobs[i] = ji
		}
	}

	for i, js := range jobStrings {
		fmt.Printf("Job %d: %s\n", i, js)
	}
}
