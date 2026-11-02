package certificate

func New(
	namespace string,
	name string,
	issuer string,
) *Certificate {
	return &Certificate{Name: name, Namespace: namespace, Issuer: issuer}
}
