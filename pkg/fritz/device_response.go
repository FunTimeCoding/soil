package fritz

type DeviceResponse struct {
	Model         string `xml:"Body>GetInfoResponse>NewModelName"`
	Software      string `xml:"Body>GetInfoResponse>NewSoftwareVersion"`
	UpTimeSeconds int    `xml:"Body>GetInfoResponse>NewUpTime"`
}
