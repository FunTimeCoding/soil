package argument

type IssueCertificate struct {
	Authority  string   `json:"authority"`
	Kind       string   `json:"kind"`
	CommonName string   `json:"common_name"`
	Host       []string `json:"host"`
	ValidDay   float64  `json:"valid_day"`
}
