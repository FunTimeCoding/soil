package gosalt

import (
	"github.com/funtimecoding/soil/pkg/salt/state_tree"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosalt/constant"
	"github.com/spf13/cobra"
)

func why(t *terminal.Terminal) *cobra.Command {
	var directories []string
	var environments []string
	result := &cobra.Command{
		Use:   "why [state]",
		Short: "Show how top.sls reaches a state",
		Long:  constant.WhyDescription,
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				state_tree.New(directories, environments).Why(arguments[0]),
			)
		},
	}
	treeFlags(result, &directories, &environments)

	return result
}
