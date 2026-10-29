package gomaintlog

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlog/constant"
	"github.com/spf13/cobra"
)

func entryFlags(o *cobra.Command) {
	o.Flags().String(constant.ActionFlag, "", "what was done")
	o.Flags().String(constant.UserFlag, "", "who did it")
	o.Flags().String(constant.SystemFlag, "", "the system it was done to")
	o.Flags().String(constant.ServiceFlag, "", "the service it was done to")
	o.Flags().String(constant.DescriptionFlag, "", "why, in detail")
	cobra.CheckErr(o.MarkFlagRequired(constant.ActionFlag))
	cobra.CheckErr(o.MarkFlagRequired(constant.UserFlag))
}
