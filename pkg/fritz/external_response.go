package fritz

type ExternalResponse struct {
	Address      string `xml:"Body>X_AVM-DE_GetExternalIPv6AddressResponse>NewExternalIPv6Address"`
	PrefixLength int    `xml:"Body>X_AVM-DE_GetExternalIPv6AddressResponse>NewPrefixLength"`
}
