package gomaintlog

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlog/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/spf13/cobra"
)

func request(o *cobra.Command) client.PostEntryRequest {
	action, e := o.Flags().GetString(constant.ActionFlag)
	cobra.CheckErr(e)
	user, f := o.Flags().GetString(constant.UserFlag)
	cobra.CheckErr(f)

	return client.PostEntryRequest{
		Action:      action,
		User:        user,
		System:      optional(o, constant.SystemFlag),
		Service:     optional(o, constant.ServiceFlag),
		Description: optional(o, constant.DescriptionFlag),
	}
}
