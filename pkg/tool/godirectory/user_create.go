package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func userCreate(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var mail, password string
	result := &cobra.Command{
		Use:   "create <account> <name> <surname>",
		Short: "Create a directory user",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			body := client.PostUserJSONRequestBody{
				Account: a[0],
				Name:    a[1],
				Surname: a[2],
			}

			if mail != "" {
				body.Mail = &mail
			}

			if password != "" {
				body.Password = &password
			}

			t.EmitValue(c.PostUser(context.Background(), body))
		},
	}
	result.Flags().StringVar(&mail, "mail", "", "Mail address")
	result.Flags().StringVar(&password, "password", "", "Initial password")

	return result
}
