package golinkace

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	s := instrument.New(constant.Identity, version)
	defer func() { s.Flush(recover()) }()
	var host string
	var port int
	var linkaceHost string
	c := command_context.New()
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
				environment.Exists(constant.DaemonInsecureEnvironment),
				linkaceHost,
				environment.Required(constant.DaemonTokenEnvironment),
			)
		},
		PersistentPostRun: func(
			m *cobra.Command,
			_ []string,
		) {
			s.RecordCommand(m.Name())
		},
	}
	o.PersistentFlags().StringVar(
		&host,
		"host",
		environment.Fallback(constant.DaemonHostEnvironment, web.Localhost),
		"golinkaced host",
	)
	o.PersistentFlags().IntVar(
		&port,
		"port",
		environment.FallbackInteger(
			constant.DaemonPortEnvironment,
			web.ListenPort,
		),
		"golinkaced port",
	)
	o.PersistentFlags().StringVar(
		&linkaceHost,
		"linkace-host",
		environment.Fallback(constant.LinkAceHostEnvironment, ""),
		"LinkAce hostname for format URLs",
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
	argument.CobraStamp(o, constant.Identity, version, gitHash, buildDate)
	errors.PanicOnError(o.Execute())
}
