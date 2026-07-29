package constant

type Ban struct {
	Substring bool
	Message   string
}

var ForbiddenImports = map[string]Ban{
	"flag": {
		Message: `use "github.com/spf13/pflag" instead of "flag"`,
	},
	"testify": {
		Substring: true,
		Message:   `use "github.com/funtimecoding/soil/pkg/assert" instead of testify`,
	},
}
