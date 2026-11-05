package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func addNetworkCommands(t *terminal.Terminal, c face.Client, o *cobra.Command) {
	o.AddCommand(listTunnelGroups(c, t))
	o.AddCommand(listTunnels(c, t))
	o.AddCommand(listTunnelTerminations(c, t))
	o.AddCommand(createTunnelGroup(c, t))
	o.AddCommand(createTunnel(c, t))
	o.AddCommand(createDeviceTunnelTermination(c, t))
	o.AddCommand(createVirtualTunnelTermination(c, t))
	o.AddCommand(listPrefixes(c, t))
	o.AddCommand(createPrefix(c, t))
	o.AddCommand(listCables(c, t))
	o.AddCommand(createCable(c, t))
}
