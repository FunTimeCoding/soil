package goband

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/spf13/cobra"
	"strings"
)

func status(c *band.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show provisioning, redirection, and consent state",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			general := c.MustGeneralSettings()
			console.Format(
				"Host: %s.%s\n",
				general.HostName,
				general.DomainName,
			)
			console.Format(
				"Provisioning: %s\n",
				band.SetupStateName(c.MustSetup().ProvisioningState),
			)
			screen := c.MustScreenSettings()
			console.Format(
				"Screen redirection: enabled=%v timeout=%dm\n",
				screen.EnabledInFirmware,
				screen.SessionTimeout,
			)
			console.Format(
				"Redirection listener: %v\n",
				c.MustRedirection().ListenerEnabled,
			)
			console.Format(
				"Consent: %s\n",
				band.ConsentName(c.MustConsent().Required),
			)
			state := c.MustPowerState()
			console.Format(
				"Power state: %s\n",
				band.PowerStateName(state.State),
			)
			names := make([]string, 0, len(state.Available))

			for _, a := range state.Available {
				names = append(names, band.PowerStateName(a))
			}

			console.Format(
				"Available power operations: %s\n",
				strings.Join(names, ", "),
			)
		},
	}
}
