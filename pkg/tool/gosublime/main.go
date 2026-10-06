package gosublime

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosublime/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	c, e := client.NewClientWithResponses(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		).String(),
		client.WithHTTPClient(web.StallClient()),
		client.WithRequestEditorFn(
			web.DeferredBearerEditor(t.Required, constant.TokenEnvironment),
		),
	)

	if e != nil {
		t.Exitf("client: %v\n", e)
	}

	x := &Context{Client: c, Terminal: t}
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(views(x))
	o.AddCommand(read(x))
	o.AddCommand(create(x))
	o.AddCommand(edit(x))
	o.AddCommand(open(x))
	o.AddCommand(save(x))
	o.AddCommand(closeView(x))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
