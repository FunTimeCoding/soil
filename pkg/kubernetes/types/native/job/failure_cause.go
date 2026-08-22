package job

import (
	"fmt"
	"k8s.io/api/batch/v1"
	core "k8s.io/api/core/v1"
)

func (j *Job) FailureCause() error {
	for _, c := range j.Raw.Status.Conditions {
		if c.Type == v1.JobFailed && c.Status == core.ConditionTrue {
			return fmt.Errorf("%s: %s", c.Reason, c.Message)
		}
	}

	return fmt.Errorf("failed pods: %d", j.Raw.Status.Failed)
}
