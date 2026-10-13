package response

type RequestStateChange struct {
	ReturnValue int `xml:"Body>RequestStateChange_OUTPUT>ReturnValue"`
}
