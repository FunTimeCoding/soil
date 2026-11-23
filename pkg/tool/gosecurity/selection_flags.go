package gosecurity

import (
	gosecuritydConstant "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/spf13/cobra"
)

func selectionFlags(o *cobra.Command) {
	o.Flags().String(
		gosecuritydConstant.ShowField,
		"",
		"coverage: supported, tracked or untracked",
	)
	o.Flags().StringArray(
		gosecuritydConstant.LabelField,
		nil,
		"a label the systems carry, key:value; repeat for several keys",
	)
	o.Flags().String(
		gosecuritydConstant.SearchField,
		"",
		"part of the host name, ignoring case",
	)
	scoreFlag(o)
}
