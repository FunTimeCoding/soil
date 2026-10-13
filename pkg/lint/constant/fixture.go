package constant

import "github.com/funtimecoding/soil/pkg/lint/types/restriction"

var FixtureRestrictions = []restriction.Restriction{
	{
		Package:   "example/fakegorm",
		Symbol:    "Open",
		AllowedIn: []string{"example/blessed"},
		Message:   "open through the blessed package",
	},
	{
		Package:   "example/fakegorm",
		Symbol:    "Shared",
		AllowedIn: []string{"example/blessed"},
		Message:   "share through the blessed package",
	},
}

var FixtureLiteralRestrictions = []restriction.Restriction{
	{
		Package:   "example/fakehttp",
		Symbol:    "Client",
		AllowedIn: []string{"example/blessed"},
		Message:   "build clients in the blessed package",
	},
}
