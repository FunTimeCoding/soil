package model_context

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func assignmentArgument(
	r mcp.CallToolRequest,
) *configuration_assignment.Assignment {
	arguments := r.GetArguments()
	given := false

	for _, k := range []string{
		constant.Roles,
		constant.Clusters,
		constant.Platforms,
		constant.Sites,
		constant.Tenants,
		constant.Tags,
	} {
		if _, okay := arguments[k]; okay {
			given = true
		}
	}

	if !given {
		return nil
	}

	result := configuration_assignment.New()
	result.Roles = r.GetStringSlice(constant.Roles, nil)
	result.Clusters = r.GetStringSlice(constant.Clusters, nil)
	result.Platforms = r.GetStringSlice(constant.Platforms, nil)
	result.Sites = r.GetStringSlice(constant.Sites, nil)
	result.Tenants = r.GetStringSlice(constant.Tenants, nil)
	result.Tags = r.GetStringSlice(constant.Tags, nil)

	return result
}
