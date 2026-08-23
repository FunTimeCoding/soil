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
	x := &Context{Client: client.NewEnvironment()}
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
	o.AddCommand(tasks(x))
	o.AddCommand(create(x))
	o.AddCommand(score(x))
	o.AddCommand(tags(x))
	o.AddCommand(statistic(x))
	o.AddCommand(cron(x))
	o.AddCommand(allocate(x))
	o.AddCommand(gear(x))
	o.AddCommand(equip(x))
	errors.PanicOnError(o.Execute())
}
