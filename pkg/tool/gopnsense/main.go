package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/argument"
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
	defer c.Close()
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(queryCommand("leases", "List DHCP leases", c.ListLeases, t))
	o.AddCommand(queryCommand("hosts", "List host entries", c.ListHosts, t))
	o.AddCommand(queryCommand("pools", "List DHCP pools", c.ListPools, t))
	o.AddCommand(queryCommand("rules", "List firewall rules", c.ListRules, t))
	o.AddCommand(
		queryCommand("aliases", "List firewall aliases", c.ListAliases, t),
	)
	o.AddCommand(
		queryCommand("source-nat", "List source NAT rules", c.ListSourceNat, t),
	)
	o.AddCommand(
		queryCommand(
			"forwards",
			"List Unbound query forwards",
			c.ListForwards,
			t,
		),
	)
	o.AddCommand(
		queryCommand(
			"blocklists",
			"List Unbound blocklists",
			c.ListBlocklists,
			t,
		),
	)
	o.AddCommand(interfaces(c, t))
	o.AddCommand(log(c, t))
	o.AddCommand(
		queryCommand("states", "Query the state table", c.FirewallStates, t),
	)
	o.AddCommand(addHost(c, t))
	o.AddCommand(setHost(c, t))
	o.AddCommand(deleteHost(c, t))
	o.AddCommand(reconfigure(c, t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	argument.CobraExecute(o)
}
