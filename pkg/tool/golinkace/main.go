package golinkace

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/constant"
	golinkacedConstant "github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	var host string
	var port int
	c := command_context.New(terminal.New(s))
	defer c.Close()
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
		PersistentPreRun: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Initialize(
				host,
				port,
				environment.Exists(golinkacedConstant.InsecureEnvironment),
				c.Terminal().Required(golinkacedConstant.TokenEnvironment),
			)
		},
	}
	o.PersistentFlags().StringVar(
		&host,
		"host",
		environment.Fallback(
			golinkacedConstant.HostEnvironment,
			webConstant.Localhost,
		),
		"golinkaced host",
	)
	o.PersistentFlags().IntVar(
		&port,
		"port",
		environment.FallbackInteger(
			golinkacedConstant.PortEnvironment,
			webConstant.ListenPort,
		),
		"golinkaced port",
	)
	o.AddCommand(links(c))
	o.AddCommand(lists(c))
	o.AddCommand(tags(c))
	o.AddCommand(search(c))
	o.AddCommand(searchList(c))
	o.AddCommand(searchTag(c))
	o.AddCommand(createLink(c))
	o.AddCommand(createList(c))
	o.AddCommand(createTag(c))
	o.AddCommand(edit(c))
	o.AddCommand(editList(c))
	o.AddCommand(editTag(c))
	o.AddCommand(editNote(c))
	o.AddCommand(appendTag(c))
	o.AddCommand(removeTag(c))
	o.AddCommand(appendList(c))
	o.AddCommand(removeList(c))
	o.AddCommand(deleteList(c))
	o.AddCommand(deleteTag(c))
	o.AddCommand(deleteNote(c))
	o.AddCommand(deleteBranch(c))
	o.AddCommand(deleteLink(c))
	o.AddCommand(notes(c))
	o.AddCommand(addNote(c))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	argument.CobraExecute(o)
}
