package constant

import "github.com/funtimecoding/soil/pkg/lint/types/restriction"

// Fixture
var FixtureRestrictions = []restriction.Restriction{
	{
		Package:   "example/fakegorm",
		Function:  "Open",
		AllowedIn: []string{"example/blessed"},
		Message:   "open through the blessed package",
	},
}
