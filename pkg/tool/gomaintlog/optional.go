package gomaintlog

import "github.com/spf13/cobra"

func optional(
	o *cobra.Command,
	name string,
) *string {
	result, e := o.Flags().GetString(name)
	cobra.CheckErr(e)

	if result == "" {
		return nil
	}

	return &result
}
