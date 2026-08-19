package argument

type ListCertificates struct {
	Authority     string  `json:"authority"`
	Kind          string  `json:"kind"`
	ExpiresBefore string  `json:"expires_before"`
	Revoked       bool    `json:"revoked"`
	Limit         float64 `json:"limit"`
}
