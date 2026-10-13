package gosecret

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	p := stall_client.New().Build()
	defer p.Close()
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(
		manifestCommand(
			t,
			constant.DecodeUsage,
			"Write a .decoded.txt beside every secret manifest",
			constant.Decode,
		),
	)
	o.AddCommand(
		manifestCommand(
			t,
			constant.CheckUsage,
			"Verify every .decoded.txt matches its secret manifest",
			constant.Check,
		),
	)
	o.AddCommand(
		manifestCommand(
			t,
			constant.EncodeUsage,
			"Write every .decoded.txt back into its secret manifest",
			constant.Encode,
		),
	)
	o.AddCommand(statusCommand(t, p.Client()))
	o.AddCommand(listCommand(t, p.Client()))
	o.AddCommand(showCommand(t, p.Client()))
	o.AddCommand(revealCommand(t, p.Client()))
	o.AddCommand(putCommand(t, p.Client()))
	o.AddCommand(migrateCommand(t, p.Client()))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
