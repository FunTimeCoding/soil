package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func New(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	addDeviceCommands(t, c, result)
	addCatalogCommands(t, c, result)
	addBuildCommands(t, c, result)
	addTagCommands(t, c, result)
	addNetworkCommands(t, c, result)

	return result
}
