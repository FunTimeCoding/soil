package band

type PowerStateResponse struct {
	State     int   `xml:"Body>PullResponse>Items>CIM_AssociatedPowerManagementService>PowerState"`
	Available []int `xml:"Body>PullResponse>Items>CIM_AssociatedPowerManagementService>AvailableRequestedPowerStates"`
}
