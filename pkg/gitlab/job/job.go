package job

import (
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"time"
)

type Job struct {
	MonitorIdentifier string
	Identifier        int64
	Name              string
	Status            string
	Stage             string
	Create            *time.Time
	Link              string
	Project           *project.Project // nil unless enriched
	Trace             string           // empty unless enriched and job failed
	concern           []string
	Raw               *gitlab.Job
}
