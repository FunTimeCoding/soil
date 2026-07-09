package restricted_call

var rules = []Rule{
	{
		Package:  "database/sql",
		Function: "Open",
		AllowedIn: []string{
			"github.com/funtimecoding/soil/pkg/relational",
			"github.com/funtimecoding/soil/pkg/relational/lite/connection",
		},
		Message: "open raw sqlite through relational/lite/connection.New and postgres through relational (see doc/ai/spec/database.md)",
	},
	{
		Package:  "gorm.io/gorm",
		Function: "Open",
		AllowedIn: []string{
			"github.com/funtimecoding/soil/pkg/relational",
			"github.com/funtimecoding/soil/pkg/relational/lite",
		},
		Message: "open databases through relational.Open, relational.NewMapper, lite.New, or lite.NewMemory (see doc/ai/spec/database.md)",
	},
}
