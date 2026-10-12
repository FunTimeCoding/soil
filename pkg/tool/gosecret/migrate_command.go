package gosecret

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/spf13/cobra"
)

func migrateCommand(t *terminal.Terminal) *cobra.Command {
	var namespace string
	var store string
	result := &cobra.Command{
		Use:   constant.MigrateUsage,
		Short: "Move a Secret manifest's values into the store and replace it with an ExternalSecret",
		Long:  "The values come from the manifest itself and are read back from the store before anything is written; a .decoded.txt that differs from the manifest stops the migration. The manifest keeps its file name and becomes an ExternalSecret of the same name, and the .decoded.txt beside it is removed.",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := Migrate(
				newInterfaceStore(interfaceClient(t)),
				arguments[0],
				namespace,
				store,
			)

			if e != nil {
				t.Exitf("%v\n", e)

				return
			}

			console.Format("path: %s\n", r.Path)
			console.Format("keys: %s\n", join.Comma(r.Keys))

			if len(r.Extra) > 0 {
				console.Format(
					"already in the store, now in the Secret too: %s\n",
					join.Comma(r.Extra),
				)
			}

			console.Format("wrote %s\n", arguments[0])

			if r.Removed != "" {
				console.Format("removed %s\n", r.Removed)
			}
		},
	}
	f := result.Flags()
	f.StringVar(
		&namespace,
		argumentConstant.Namespace,
		"",
		"Namespace part of the store path; default the manifest's directory name",
	)
	f.StringVar(
		&store,
		argumentConstant.Store,
		constant.DefaultStore,
		"ClusterSecretStore the ExternalSecret reads from",
	)

	return result
}
