package gosecurity

import (
	gosecuritydConstant "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/spf13/cobra"
)

func score(o *cobra.Command) *float64 {
	if !o.Flags().Changed(gosecuritydConstant.ScoreField) {
		return nil
	}

	result, e := o.Flags().GetFloat64(gosecuritydConstant.ScoreField)
	cobra.CheckErr(e)

	return &result
}
