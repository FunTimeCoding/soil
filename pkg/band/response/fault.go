package response

type Fault struct {
	Text string `xml:"Body>Fault>Reason>Text"`
}
