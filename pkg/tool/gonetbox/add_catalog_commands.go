package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func addCatalogCommands(t *terminal.Terminal, c *client.Client, o *cobra.Command) {
	o.AddCommand(listAddressRanges(c, t))
	o.AddCommand(createAddressRange(c, t))
	o.AddCommand(listLocations(c, t))
	o.AddCommand(createLocation(c, t))
	o.AddCommand(listPlatforms(c, t))
	o.AddCommand(createPlatform(c, t))
	o.AddCommand(listConfigurationContexts(c, t))
	o.AddCommand(getConfigurationContext(c, t))
	o.AddCommand(createConfigurationContext(c, t))
	o.AddCommand(updateConfigurationContext(c, t))
	o.AddCommand(deleteConfigurationContext(c, t))
	o.AddCommand(ansibleInventory(c, t))
	o.AddCommand(listChoiceSets(c, t))
	o.AddCommand(getChoiceSet(c, t))
	o.AddCommand(createChoiceSet(c, t))
	o.AddCommand(updateChoiceSet(c, t))
	o.AddCommand(deleteChoiceSet(c, t))
	o.AddCommand(listCustomFields(c, t))
	o.AddCommand(createCustomField(c, t))
	o.AddCommand(deleteCustomField(c, t))
}
