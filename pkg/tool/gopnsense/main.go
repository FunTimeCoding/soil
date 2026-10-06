package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsense/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(queryCommand("leases", "List DHCP leases", c.Leases, t))
	o.AddCommand(queryCommand("hosts", "List host entries", c.Hosts, t))
	o.AddCommand(queryCommand("pools", "List DHCP pools", c.Pools, t))
	o.AddCommand(queryCommand("rules", "List firewall rules", c.Rules, t))
	o.AddCommand(queryCommand("aliases", "List firewall aliases", c.Aliases, t))
	o.AddCommand(
		queryCommand("source-nat", "List source NAT rules", c.SourceNat, t),
	)
	o.AddCommand(
		queryCommand("forwards", "List Unbound query forwards", c.Forwards, t),
	)
	o.AddCommand(
		queryCommand("blocklists", "List Unbound blocklists", c.Blocklists, t),
	)
	o.AddCommand(interfaces(c, t))
	o.AddCommand(log(c, t))
	o.AddCommand(queryCommand("states", "Query the state table", c.States, t))
	o.AddCommand(addHost(c, t))
	o.AddCommand(setHost(c, t))
	o.AddCommand(deleteHost(c, t))
	o.AddCommand(reconfigure(c, t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
