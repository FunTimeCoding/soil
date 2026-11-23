package gosalt

import (
	"github.com/funtimecoding/soil/pkg/salt/state_tree"
	"github.com/funtimecoding/soil/pkg/salt/types/orphan_report"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosalt/constant"
	"github.com/spf13/cobra"
)

func orphans(t *terminal.Terminal) *cobra.Command {
	var directories []string
	var environments []string
	var pillar []string
	result := &cobra.Command{
		Use:   "orphans",
		Short: "Report unreached states and files",
		Long:  constant.OrphansDescription,
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			s := state_tree.New(directories, environments)
			s.Cite(state_tree.New(pillar, environments).Served())
			t.EmitValue(orphan_report.New(s.Analyze()), nil)
		},
	}
	treeFlags(result, &directories, &environments)
	result.Flags().StringSliceVar(
		&pillar,
		constant.PillarArgument,
		nil,
		constant.PillarUsage,
	)

	return result
}
