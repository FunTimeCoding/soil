package constant

import "github.com/funtimecoding/soil/pkg/lint/types/ban"

var ForbiddenImports = map[string]ban.Ban{
	"flag": {Message: `use "github.com/spf13/pflag" instead of "flag"`},
	"testify": {
		Substring: true,
		Message:   `use "github.com/funtimecoding/soil/pkg/assert" instead of testify`,
	},
}

var ForbiddenCalls = map[string]bool{
	"Command":        true,
	"CommandContext": true,
}
