package changed

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func Upstream(directory string) string {
	for _, reference := range []string{constant.Upstream, constant.RemoteHead} {
		r := run.New()
		r.Panic = false
		r.Directory = directory
		r.Start(
			constant.Command,
			constant.RevParse,
			constant.AbbreviatedReference,
			constant.SymbolicFullName,
			reference,
		)

		if r.Error == nil {
			if s := strings.TrimSpace(r.OutputString); s != "" {
				return s
			}
		}
	}

	return ""
}
