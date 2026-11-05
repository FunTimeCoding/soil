package goquery

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
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
	result.AddCommand(search(c, t))
	result.AddCommand(list(c, t))
	result.AddCommand(index(c, t))
	result.AddCommand(addCollection(c, t))
	result.AddCommand(deleteCollection(c, t))
	result.AddCommand(embed(c, t))
	result.AddCommand(addContext(c, t))
	result.AddCommand(listContexts(c, t))
	result.AddCommand(listTags(c, t))
	result.AddCommand(listMetadata(c, t))
	result.AddCommand(removeContext(c, t))
	result.AddCommand(status(c, t))
	result.AddCommand(oversize(c, t))
	result.AddCommand(failed(c, t))
	result.AddCommand(chunk(c, t))
	result.AddCommand(rechunk(c, t))

	return result
}
