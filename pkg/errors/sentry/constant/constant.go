package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

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

	Response = "response"
	Body     = "body"
)

const (
	PeriodEmpty     = ""
	PeriodDay       = "24h"
	PeriodFortnight = "14d"

	SortNewestFirst = "-timestamp"
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

	Periods = []string{
		PeriodEmpty,
		PeriodDay,
		PeriodFortnight,
	}

	UnresolvedFilter = "is:unresolved"

	ErrorNotFound = errors.New("not found")
	Format        = option.Color.Copy()
)
