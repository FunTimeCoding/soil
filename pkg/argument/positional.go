package argument

import "github.com/spf13/pflag"

func Positional(number int) string {
	return pflag.Arg(number)
}
