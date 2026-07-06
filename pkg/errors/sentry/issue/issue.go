package issue

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/face"
	"time"
)

type Issue struct {
	Project           string
	Type              string
	Title             string
	Link              string
	Create            *time.Time
	MonitorIdentifier string
	ageColor          face.SprintFunction
	Raw               *response.Issue
}
