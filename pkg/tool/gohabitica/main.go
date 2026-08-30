package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/gohabitica/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	s := instrument.New(constant.Identity, version)
	defer func() { s.Flush(recover()) }()
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
		PersistentPostRun: func(
			m *cobra.Command,
			_ []string,
		) {
			s.RecordCommand(m.Name())
		},
	}
	o.AddCommand(tasks(c))
	o.AddCommand(create(c))
	o.AddCommand(score(c))
	o.AddCommand(tags(c))
	o.AddCommand(statistic(c))
	o.AddCommand(cron(c))
	o.AddCommand(allocate(c))
	o.AddCommand(gear(c))
	o.AddCommand(equip(c))
	errors.PanicOnError(o.Execute())
}
