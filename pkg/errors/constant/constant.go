package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/console/constant"
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

	Response        = "response"
	Body            = "body"
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

	Periods = []string{
		PeriodEmpty,
		PeriodDay,
		PeriodFortnight,
	}
)

var (
	ErrorNotFound = errors.New("not found")
	Format        = constant.ColorFormat.Copy()
)
