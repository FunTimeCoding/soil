package goprocess

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goprocess/constant"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	o := &cobra.Command{
		Use:           constant.Identity.Usage(),
		Short:         constant.Identity.Description(),
		Version:       argument.CobraVersion(version, gitHash, buildDate),
		SilenceErrors: true,
	}
	o.PersistentFlags().StringP(
		"file",
		"f",
		"Procfile",
		"path to Procfile",
	)
	o.AddCommand(checkCommand())
	o.AddCommand(runCommand())
	errors.PanicOnError(o.Execute())
}
