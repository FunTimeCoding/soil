package response

type CertificateSummary struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	DnsNames  string `json:"dnsNames"`
	Ready     string `json:"ready"`
	Expires   string `json:"expires"`
	ExpiresAt string `json:"expiresAt"`
	Renewal   string `json:"renewal,omitempty"`
}
