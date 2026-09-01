package client

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/job"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"time"
)

func (c *Client) WaitForJob(
	namespace string,
	name string,
	limit time.Duration,
) error {
	start := time.Now()

	for {
		if limit > 0 && time.Since(start) > limit {
			return timeout.Format("job timeout: %s", name)
		}

		j := c.Job(namespace, name)

		if j == nil {
			return nil
		}

		if j.Raw.Status.CompletionTime != nil {
			console.Format("job done: %s\n", name)

			return nil
		}

		if j.Raw.Status.Failed > 0 {
			failure := job.New(name, "kubernetes", j.FailureCause())
			failure.Detail = map[string]any{"namespace": namespace}

			return failure
		}

		console.Format("job running: %s\n", name)
		time.Sleep(10 * time.Second)
	}
}
