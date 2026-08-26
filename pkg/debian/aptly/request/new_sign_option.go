package request

func NewSignOption(passphraseFile string) *SignOption {
	return &SignOption{PassphraseFile: passphraseFile}
}
