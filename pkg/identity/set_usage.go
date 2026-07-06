package identity

import (
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/spf13/pflag"
	"os"
)

func (t *Tool) SetUsage() {
	pflag.Usage = func() {
		writer.Print(
			os.Stderr,
			"%s - %s\n\nUsage:\n  %s\n\nFlags:\n",
			t.name,
			t.description,
			t.usage,
		)
		pflag.PrintDefaults()
	}
}
