package band

import (
	"encoding/xml"
	"fmt"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
)

func (c *Client) RequestPowerState(state int) error {
	selectors := `<w:SelectorSet><w:Selector Name="CreationClassName">CIM_PowerManagementService</w:Selector><w:Selector Name="Name">Intel(r) AMT Power Management Service</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet>`
	inner := fmt.Sprintf(
		`<p:RequestPowerStateChange_INPUT xmlns:p="%s"><p:PowerState>%d</p:PowerState><p:ManagedElement><a:Address>%s</a:Address><a:ReferenceParameters><w:ResourceURI>%s</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="Name">ManagedSystem</w:Selector></w:SelectorSet></a:ReferenceParameters></p:ManagedElement></p:RequestPowerStateChange_INPUT>`,
		constant.PowerServiceResource,
		state,
		constant.AnonymousRole,
		constant.ComputerSystemResource,
	)
	body, e := c.call(
		constant.PowerStateAction,
		constant.PowerServiceResource,
		selectors,
		inner,
	)

	if e != nil {
		return e
	}

	var result response.RequestPowerState

	if f := xml.Unmarshal(body, &result); f != nil {
		return f
	}

	if result.ReturnValue != 0 {
		return unexpected.Format(
			"power state change %d refused: return value %d",
			state,
			result.ReturnValue,
		)
	}

	return nil
}
