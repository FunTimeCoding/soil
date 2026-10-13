package constant

import "github.com/funtimecoding/soil/pkg/lint/types/restriction"

var RestrictedCalls = []restriction.Restriction{
	{
		Package: "database/sql",
		Symbol:  "Open",
		AllowedIn: []string{
			"github.com/funtimecoding/soil/pkg/relational",
			"github.com/funtimecoding/soil/pkg/relational/lite/connection",
		},
		Message: "open raw sqlite through relational/lite/connection.New and postgres through relational (see doc/ai/spec/database.md)",
	},
	{
		Package: "gorm.io/gorm",
		Symbol:  "Open",
		AllowedIn: []string{
			"github.com/funtimecoding/soil/pkg/relational",
			"github.com/funtimecoding/soil/pkg/relational/lite",
		},
		Message: "open databases through relational.Open, relational.NewMapper, lite.New, or lite.NewMemory (see doc/ai/spec/database.md)",
	},
	{Package: HTTPPackage, Symbol: "Get", Message: OwnedClientMessage},
	{Package: HTTPPackage, Symbol: "Post", Message: OwnedClientMessage},
	{Package: HTTPPackage, Symbol: "Head", Message: OwnedClientMessage},
	{Package: HTTPPackage, Symbol: "PostForm", Message: OwnedClientMessage},
	{
		Package: HTTPPackage,
		Symbol:  "DefaultClient",
		Message: OwnedClientMessage,
	},
	{
		Package: HTTPPackage,
		Symbol:  "DefaultTransport",
		Message: OwnedClientMessage,
	},
}
