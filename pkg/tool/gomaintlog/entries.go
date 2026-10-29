package gomaintlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlog/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/spf13/cobra"
)

func entries(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "entries",
		Short: "List maintenance log entries, newest first",
		Args:  cobra.NoArgs,
		Run: func(
			o *cobra.Command,
			_ []string,
		) {
			p := &client.GetEntriesParams{
				System:  optional(o, constant.SystemFlag),
				Service: optional(o, constant.ServiceFlag),
				User:    optional(o, constant.UserFlag),
			}
			limit, e := o.Flags().GetInt(constant.LimitFlag)
			cobra.CheckErr(e)

			if limit > 0 {
				p.Limit = &limit
			}

			t.EmitValue(c.GetEntries(context.Background(), p))
		},
	}
	result.Flags().String(
		constant.SystemFlag,
		"",
		"only entries for this system",
	)
	result.Flags().String(
		constant.ServiceFlag,
		"",
		"only entries for this service",
	)
	result.Flags().String(constant.UserFlag, "", "only entries by this user")
	result.Flags().Int(constant.LimitFlag, 0, "at most this many entries")

	return result
}
