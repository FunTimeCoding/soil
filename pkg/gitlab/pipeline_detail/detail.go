package pipeline_detail

import (
	"gitlab.com/gitlab-org/api/client-go/v2"
	"time"
)

type Detail struct {
	Identifier        int64
	ProjectIdentifier int64
	Status            string
	Source            string
	Reference         string
	Hash              string
	Link              string
	Create            *time.Time
	Update            *time.Time
	Raw               *gitlab.Pipeline
}
