package band

type RedirectionResponse struct {
	EnabledState    int  `xml:"Body>AMT_RedirectionService>EnabledState"`
	ListenerEnabled bool `xml:"Body>AMT_RedirectionService>ListenerEnabled"`
}
