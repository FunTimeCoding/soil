package argument_tester

import "github.com/spf13/cobra"

func Root(
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
