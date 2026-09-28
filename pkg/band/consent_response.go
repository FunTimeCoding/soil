package band

type ConsentResponse struct {
	Required int `xml:"Body>IPS_OptInService>OptInRequired"`
}
