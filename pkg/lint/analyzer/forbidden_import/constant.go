package forbidden_import

type ban struct {
	substring bool
	message   string
}

var banned = map[string]ban{
	"flag": {
		message: `use "github.com/spf13/pflag" instead of "flag"`,
	},
	"testify": {
		substring: true,
		message:   `use "github.com/funtimecoding/soil/pkg/assert" instead of testify`,
	},
}
