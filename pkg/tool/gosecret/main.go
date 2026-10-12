package gosecret

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
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
	o.AddCommand(statusCommand(t))
	o.AddCommand(listCommand(t))
	o.AddCommand(showCommand(t))
	o.AddCommand(revealCommand(t))
	o.AddCommand(putCommand(t))
	o.AddCommand(migrateCommand(t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
