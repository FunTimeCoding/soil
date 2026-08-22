package constant

import "github.com/funtimecoding/soil/pkg/console/constant"

const (
	Base = "/api/0"

	UndefinedEnvironment = "undefined"
	UndefinedVersion     = "undefined"

	HostEnvironment         = "SENTRY_HOST"
	TokenEnvironment        = "SENTRY_TOKEN"
	OrganizationEnvironment = "SENTRY_ORGANIZATION"
	ProjectEnvironment      = "SENTRY_PROJECT"
	LocatorEnvironment      = "SENTRY_LOCATOR"

	ErrorType = "error"

	Response        = "response"
	Body            = "body"
	Process         = "process"
	Command         = "command"
	Output          = "output"
	Stderr          = "stderr"
	Job             = "job"
	Identifier      = "identifier"
	Kind            = "kind"
	Upstream        = "upstream"
	Status          = "status"
	Detail          = "detail"
	PeriodEmpty     = ""
	PeriodDay       = "24h"
	PeriodFortnight = "14d"

	SortNewestFirst = "-timestamp"

	UnresolvedFilter = "is:unresolved"
)

var (
	EventFields = []string{
		"id",
		"title",
		"message",
		"project",
		"timestamp",
		"culprit",
	}

	Periods = []string{PeriodEmpty, PeriodDay, PeriodFortnight}
)

var (
	Format = constant.ColorFormat.Copy()
)
