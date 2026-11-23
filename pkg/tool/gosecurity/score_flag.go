package gosecurity

import (
	gosecuritydConstant "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/spf13/cobra"
)

func scoreFlag(o *cobra.Command) {
	o.Flags().Float64(
		gosecuritydConstant.ScoreField,
		0,
		"the lowest score a finding needs, e.g. 9 for critical",
	)
}
