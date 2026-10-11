package gotechnitium

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/technitium/response"
)

func printZoneOptions(o *response.ZoneOptions) {
	console.Format("update: %s\n", o.Update)
	console.Format("update network: %s\n", join.Comma(o.UpdateNetwork))

	for _, p := range o.UpdatePolicies {
		console.Format(
			"update policy: %s %s %s\n",
			p.KeyName,
			p.Domain,
			join.Comma(p.Types),
		)
	}

	console.Format("zone transfer: %s\n", o.ZoneTransfer)
	console.Format(
		"zone transfer network: %s\n",
		join.Comma(o.ZoneTransferNetwork),
	)
	console.Format(
		"zone transfer key: %s\n",
		join.Comma(o.ZoneTransferKeyNames),
	)
}
