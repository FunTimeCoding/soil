package response

type RequestPowerState struct {
	ReturnValue int `xml:"Body>RequestPowerStateChange_OUTPUT>ReturnValue"`
}
