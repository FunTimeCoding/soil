package certificate

type Certificate struct {
	Name      string
	Namespace string
	Ready     bool
	Issuing   bool
	Issuer    string
}
