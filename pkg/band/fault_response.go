package band

type FaultResponse struct {
	Text string `xml:"Body>Fault>Reason>Text"`
}
