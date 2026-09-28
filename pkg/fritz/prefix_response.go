package fritz

type PrefixResponse struct {
	Prefix       string `xml:"Body>X_AVM-DE_GetIPv6PrefixResponse>NewIPv6Prefix"`
	PrefixLength int    `xml:"Body>X_AVM-DE_GetIPv6PrefixResponse>NewPrefixLength"`
}
