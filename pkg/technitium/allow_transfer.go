package technitium

import (
	"github.com/funtimecoding/soil/pkg/technitium/response"
	"slices"
)

func AllowTransfer(
	o *response.ZoneOptions,
	network string,
	keyName string,
) {
	o.ZoneTransfer = WidenAccess(o.ZoneTransfer)

	if !slices.Contains(o.ZoneTransferNetwork, network) {
		o.ZoneTransferNetwork = append(o.ZoneTransferNetwork, network)
	}

	if !slices.Contains(o.ZoneTransferKeyNames, keyName) {
		o.ZoneTransferKeyNames = append(o.ZoneTransferKeyNames, keyName)
	}
}
