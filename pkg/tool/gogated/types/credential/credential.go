package credential

type Credential struct {
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
	Notice     string `json:"notice"`
}
