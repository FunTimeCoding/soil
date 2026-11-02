package response

type Fault struct {
	Text    string `xml:"Body>Fault>Reason>Text"`
	Subcode string `xml:"Body>Fault>Code>Subcode>Value"`
}
