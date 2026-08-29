package gopnsense

import "github.com/spf13/cobra"

func registerHostFlags(
	c *cobra.Command,
	f *hostFlags,
) {
	c.Flags().StringVar(&f.host, "host", "", "hostname without the domain")
	c.Flags().StringVar(&f.domain, "domain", "", "domain of the host")
	c.Flags().StringVar(&f.address, "address", "", "address to map or reserve")
	c.Flags().StringVar(
		&f.hardwareAddress,
		"hardware-address",
		"",
		"MAC address the reservation binds to",
	)
	c.Flags().StringVar(
		&f.clientIdentifier,
		"client-identifier",
		"",
		"DHCP client identifier",
	)
	c.Flags().StringVar(
		&f.description,
		"description",
		"",
		"description of the entry",
	)
	c.Flags().BoolVar(
		&f.apply,
		"apply",
		true,
		"reconfigure Dnsmasq after the write",
	)
}
