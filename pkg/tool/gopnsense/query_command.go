package gopnsense

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/spf13/cobra"
)

func queryCommand[T any](
	use string,
	short string,
	call func(
		x context.Context,
		phrase string,
	) ([]T, error),
	t *terminal.Terminal,
) *cobra.Command {
	var phrase string
	result := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(call(context.Background(), phrase))
		},
	}
	result.Flags().StringVar(&phrase, "query", "", "search phrase")

	return result
}
