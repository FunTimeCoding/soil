package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gopnsense/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(queryCommand("leases", "List DHCP leases", c.Leases))
	o.AddCommand(queryCommand("hosts", "List host entries", c.Hosts))
	o.AddCommand(queryCommand("pools", "List DHCP pools", c.Pools))
	o.AddCommand(queryCommand("rules", "List firewall rules", c.Rules))
	o.AddCommand(queryCommand("aliases", "List firewall aliases", c.Aliases))
	o.AddCommand(
		queryCommand("source-nat", "List source NAT rules", c.SourceNat),
	)
	o.AddCommand(
		queryCommand("forwards", "List Unbound query forwards", c.Forwards),
	)
	o.AddCommand(
		queryCommand("blocklists", "List Unbound blocklists", c.Blocklists),
	)
	o.AddCommand(interfaces(c))
	o.AddCommand(log(c))
	o.AddCommand(queryCommand("states", "Query the state table", c.States))
	errors.PanicOnError(o.Execute())
}
