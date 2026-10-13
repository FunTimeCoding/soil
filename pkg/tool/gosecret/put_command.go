package gosecret

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"strings"
)

func putCommand(
	t *terminal.Terminal,
	h *http.Client,
) *cobra.Command {
	return &cobra.Command{
		Use:   constant.PutUsage,
		Short: "Set one key from standard input, merged into the secret",
		Long:  "The value is read from standard input so it never appears in shell history or the process list. A trailing newline is dropped.",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			b, e := io.ReadAll(os.Stdin)
			errors.PanicOnError(e)
			r, f := interfaceClient(t, h).PutSecretWithResponse(
				context.Background(),
				&client.PutSecretParams{Path: arguments[0]},
				client.PutSecretJSONRequestBody{
					Values: map[string]string{
						arguments[1]: strings.TrimSuffix(string(b), "\n"),
					},
				},
			)
			errors.PanicOnError(f)

			if r.JSON200 == nil {
				t.Reject(r.Status(), r.Body)

				return
			}

			for _, k := range r.JSON200.Keys {
				console.Line(k)
			}
		},
	}
}
