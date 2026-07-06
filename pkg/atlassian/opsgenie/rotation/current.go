package rotation

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/compact"
	"time"
)

func (r *Rotation) current() string {
	weeks := int(time.Since(*r.Start).Hours() / 24 / 7)
	index := weeks % len(r.Participants)

	return compact.Mail(r.Participants[index].Username)
}
