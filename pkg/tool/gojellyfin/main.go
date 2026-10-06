package gojellyfin

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfin/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
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

	x := &Context{Client: c, Telemetry: s.Recorder(), Terminal: t}
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(libraries(x))
	o.AddCommand(search(x))
	o.AddCommand(item(x))
	o.AddCommand(episodes(x))
	o.AddCommand(tracks(x))
	o.AddCommand(sessions(x))
	o.AddCommand(play(x))
	o.AddCommand(command(x))
	o.AddCommand(volume(x))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
