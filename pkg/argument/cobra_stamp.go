package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/stamp/report"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/spf13/cobra"
)

func CobraStamp(
	c *cobra.Command,
	t *identity.Tool,
) {
	s := stamp.New()
	asNotation := c.PersistentFlags().Bool(
		constant.Notation,
		false,
		constant.NotationUsage,
	)
	c.Version = s.DisplayVersion()
	cobra.AddTemplateFunc(
		constant.StampFunction,
		func() string {
			if *asNotation {
				return join.Empty(
					notation.MarshalIndent(report.New(t.Name(), s)),
					stringsConstant.Unix,
				)
			}

			return s.Text()
		},
	)
	c.SetVersionTemplate(constant.StampTemplate)
}
