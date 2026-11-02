package gonetbox

import (
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func assignmentFlags(c *cobra.Command) func() *client.ConfigurationAssignment {
	var roles []string
	var clusters []string
	var platforms []string
	var sites []string
	var tenants []string
	var tags []string
	f := c.Flags()
	f.StringSliceVar(
		&roles,
		gonetboxdConstant.Role,
		nil,
		"role names it applies to",
	)
	f.StringSliceVar(
		&clusters,
		gonetboxdConstant.Cluster,
		nil,
		"cluster names it applies to",
	)
	f.StringSliceVar(
		&platforms,
		gonetboxdConstant.Platform,
		nil,
		"platform names it applies to",
	)
	f.StringSliceVar(
		&sites,
		gonetboxdConstant.Site,
		nil,
		"site names it applies to",
	)
	f.StringSliceVar(
		&tenants,
		gonetboxdConstant.Tenant,
		nil,
		"tenant names it applies to",
	)
	f.StringSliceVar(
		&tags,
		gonetboxdConstant.Tag,
		nil,
		"tag names it applies to",
	)

	return func() *client.ConfigurationAssignment {
		given := false

		for _, n := range []string{
			gonetboxdConstant.Role,
			gonetboxdConstant.Cluster,
			gonetboxdConstant.Platform,
			gonetboxdConstant.Site,
			gonetboxdConstant.Tenant,
			gonetboxdConstant.Tag,
		} {
			if f.Changed(n) {
				given = true
			}
		}

		if !given {
			return nil
		}

		return &client.ConfigurationAssignment{
			Roles:     &roles,
			Clusters:  &clusters,
			Platforms: &platforms,
			Sites:     &sites,
			Tenants:   &tenants,
			Tags:      &tags,
		}
	}
}
