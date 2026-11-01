package goaptly

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goaptly/constant"
	"github.com/spf13/cobra"
)

func publish(
	c *aptly.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "publish",
		Short: "Re-publish a distribution, signing it",
		Args:  cobra.NoArgs,
		Run: func(
			o *cobra.Command,
			_ []string,
		) {
			distribution, e := o.Flags().GetString(constant.DistributionFlag)
			cobra.CheckErr(e)
			passphrase, f := o.Flags().GetString(constant.PassphraseFileFlag)
			cobra.CheckErr(f)
			t.EmitDone(c.UpdatePublish(distribution, passphrase))
		},
	}
	result.Flags().String(
		constant.DistributionFlag,
		constant.DefaultRepository,
		"Published distribution to update",
	)
	result.Flags().String(
		constant.PassphraseFileFlag,
		"",
		"Signing passphrase file, as a path on the aptly server",
	)
	errors.PanicOnError(result.MarkFlagRequired(constant.PassphraseFileFlag))

	return result
}
