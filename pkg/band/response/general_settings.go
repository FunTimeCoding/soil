package response

type GeneralSettings struct {
	HostName   string `xml:"Body>AMT_GeneralSettings>HostName"`
	DomainName string `xml:"Body>AMT_GeneralSettings>DomainName"`
}
