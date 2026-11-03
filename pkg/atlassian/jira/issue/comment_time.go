package issue

import (
	"github.com/andygrunwald/go-jira"
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func CommentTime(c *jira.Comment) time.Time {
	return soilTime.PickLatest(ConvertTime(c.Created), ConvertTime(c.Updated))
}
