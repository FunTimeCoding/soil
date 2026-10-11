package request

type SignOption struct {
	Skip           bool   `json:"Skip"`
	Batch          bool   `json:"Batch,omitzero"`
	GpgKey         string `json:"GpgKey,omitempty"`
	Keyring        string `json:"Keyring,omitempty"`
	SecretKeyring  string `json:"SecretKeyring,omitempty"`
	Passphrase     string `json:"Passphrase,omitempty"`
	PassphraseFile string `json:"PassphraseFile,omitempty"`
}
