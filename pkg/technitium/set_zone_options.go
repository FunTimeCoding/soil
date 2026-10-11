package technitium

import (
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/policy"
	"github.com/funtimecoding/soil/pkg/technitium/response"
	"net/url"
)

func (c *Client) SetZoneOptions(o *response.ZoneOptions) error {
	return c.post(
		constant.ZoneOptionsSetPath,
		url.Values{
			constant.ZoneParameter:          {o.Name},
			constant.UpdateParameter:        {o.Update},
			constant.UpdateNetworkParameter: {listValue(o.UpdateNetwork)},
			constant.UpdatePolicyParameter:  {policy.Encode(o.UpdatePolicies)},
			constant.ZoneTransferParameter:  {o.ZoneTransfer},
			constant.ZoneTransferNetworkParameter: {
				listValue(o.ZoneTransferNetwork),
			},
			constant.ZoneTransferKeyParameter: {
				listValue(o.ZoneTransferKeyNames),
			},
		},
	)
}
