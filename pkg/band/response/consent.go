package response

type Consent struct {
	Required int `xml:"Body>IPS_OptInService>OptInRequired"`
}
