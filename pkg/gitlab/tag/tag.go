package tag

import (
	"gitlab.com/gitlab-org/api/client-go/v2"
	"time"
)

type Tag struct {
	Name      string
	Message   string
	Target    string
	Hash      string
	Protected bool
	Create    *time.Time
	Raw       *gitlab.Tag
}
