package gosublime

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gosublime/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	s := instrument.New(constant.Identity, version)
	defer func() { s.Flush(recover()) }()
	host := environment.Optional(constant.HostEnvironment)

	if host == "" {
		host = constant.DefaultHost
	}

	c, e := client.NewClientWithResponses(
		locator.New(host).Insecure().String(),
		client.WithRequestEditorFn(
			web.BearerEditor(environment.Required(constant.TokenEnvironment)),
		),
	)

	if e != nil {
		errors.Printf("client: %v\n", e)
		os.Exit(1)
	}

	x := &Context{Client: c}
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
	o.AddCommand(views(x))
	o.AddCommand(read(x))
	o.AddCommand(create(x))
	o.AddCommand(edit(x))
	o.AddCommand(open(x))
	o.AddCommand(save(x))
	o.AddCommand(closeView(x))
	errors.PanicOnError(o.Execute())
}
