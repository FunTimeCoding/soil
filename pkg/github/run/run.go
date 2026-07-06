package run

import (
	"github.com/funtimecoding/soil/pkg/github/job"
	"github.com/funtimecoding/soil/pkg/github/repository"
	"github.com/google/go-github/v88/github"
	"time"
)

type Run struct {
	MonitorIdentifier string
	Identifier        int64
	Name              string
	Status            string
	Conclusion        string
	Branch            string
	Create            time.Time
	Update            time.Time
	concern           []string
	Repository        *repository.Repository
	Jobs              []*job.Job
	Raw               *github.WorkflowRun
}
