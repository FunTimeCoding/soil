package constant

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/detail"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

const (
	UserKeyEnvironment = "OPSGENIE_USER_KEY"
	TeamKeyEnvironment = "OPSGENIE_TEAM_KEY"
	WebHostEnvironment = "OPSGENIE_WEB_HOST"
	TeamEnvironment    = "OPSGENIE_TEAM"

	PageLimit int = 100

	NoKey = "no key"
)

var Format = option.ExtendedColor.Copy().Tag(
	tag.Category,
)

type ParseDescription func(string) *detail.Prometheus
