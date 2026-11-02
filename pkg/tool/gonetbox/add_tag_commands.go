package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func addTagCommands(t *terminal.Terminal, c *client.Client, o *cobra.Command) {
	o.AddCommand(listTags(c, t))
	o.AddCommand(createTag(c, t))
	o.AddCommand(addDeviceTag(c, t))
	o.AddCommand(removeDeviceTag(c, t))
	o.AddCommand(addVirtualTag(c, t))
	o.AddCommand(removeVirtualTag(c, t))
	o.AddCommand(addDeviceJournalEntry(c, t))
	o.AddCommand(listDeviceJournalEntries(c, t))
	o.AddCommand(addVirtualJournalEntry(c, t))
	o.AddCommand(listVirtualJournalEntries(c, t))
	o.AddCommand(updateJournalEntry(c, t))
	o.AddCommand(deleteJournalEntry(c, t))
	o.AddCommand(deleteAddress(c, t))
	o.AddCommand(setDeviceLabel(c, t))
	o.AddCommand(removeDeviceLabel(c, t))
	o.AddCommand(listDeviceLabels(c, t))
	o.AddCommand(setVirtualLabel(c, t))
	o.AddCommand(removeVirtualLabel(c, t))
	o.AddCommand(listVirtualLabels(c, t))
}
