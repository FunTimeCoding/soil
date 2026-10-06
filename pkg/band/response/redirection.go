package response

type Redirection struct {
	EnabledState    int  `xml:"Body>AMT_RedirectionService>EnabledState"`
	ListenerEnabled bool `xml:"Body>AMT_RedirectionService>ListenerEnabled"`
}
