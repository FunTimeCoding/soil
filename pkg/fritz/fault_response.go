package fritz

type FaultResponse struct {
	Code        string `xml:"Body>Fault>detail>UPnPError>errorCode"`
	Description string `xml:"Body>Fault>detail>UPnPError>errorDescription"`
}
