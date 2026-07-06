package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createCluster(c *client.Client) *cobra.Command {
	var clusterType string
	var site string
	result := &cobra.Command{
		Use:   "create-cluster [name]",
		Short: "Create a NetBox cluster",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateCluster(arguments[0], clusterType, site))
		},
	}
	result.Flags().StringVar(
		&clusterType,
		"type",
		"",
		"cluster type name (required)",
	)
	result.Flags().StringVar(
		&site,
		"site",
		"",
		"site name (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("type"))
	errors.PanicOnError(result.MarkFlagRequired("site"))

	return result
}
