package integration

import "github.com/spf13/cobra"

func cobraRoot(
	arguments []string,
	failure error,
) *cobra.Command {
	result := &cobra.Command{Use: "alfa"}
	c := &cobra.Command{
		Use: "bravo",
		RunE: func(
			*cobra.Command,
			[]string,
		) error {
			return failure
		},
	}
	c.Flags().Bool("charlie", false, "charlie flag")
	result.AddCommand(c)
	result.SetArgs(arguments)

	return result
}
