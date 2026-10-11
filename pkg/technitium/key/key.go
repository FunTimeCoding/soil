package key

type Key struct {
	Name      string `json:"keyName"`
	Secret    string `json:"sharedSecret"`
	Algorithm string `json:"algorithmName"`
}
