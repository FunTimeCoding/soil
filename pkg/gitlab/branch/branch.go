package branch

import (
	"github.com/funtimecoding/soil/pkg/face"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"time"
)

type Branch struct {
	Name       string
	Merged     bool
	CommitDate *time.Time
	ageColor   face.SprintFunction
	Raw        *gitlab.Branch
}
