package note

import (
	"gitlab.com/gitlab-org/api/client-go/v2"
	"time"
)

type Note struct {
	Identifier int64
	Author     string
	Body       string
	System     bool
	Create     *time.Time
	Raw        *gitlab.Note
}
