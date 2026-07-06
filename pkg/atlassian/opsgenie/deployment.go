package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/deployment"
)

func (c *Client) Deployment(identifier string) *deployment.GetDeploymentResult {
	result, e := c.userClient.Deployment.Get(
		c.context,
		&deployment.GetDeploymentRequest{
			IdentifierType:  deployment.DEPLOYMENT_ID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return result
}
