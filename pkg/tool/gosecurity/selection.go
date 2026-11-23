package gosecurity

import (
	gosecuritydConstant "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/spf13/cobra"
)

func selection(o *cobra.Command) *client.ListSystemsParams {
	result := &client.ListSystemsParams{Score: score(o)}
	show, e := o.Flags().GetString(gosecuritydConstant.ShowField)
	cobra.CheckErr(e)

	if show != "" {
		result.Show = &show
	}

	labels, f := o.Flags().GetStringArray(gosecuritydConstant.LabelField)
	cobra.CheckErr(f)

	if len(labels) > 0 {
		result.Label = &labels
	}

	search, g := o.Flags().GetString(gosecuritydConstant.SearchField)
	cobra.CheckErr(g)

	if search != "" {
		result.Search = &search
	}

	return result
}
