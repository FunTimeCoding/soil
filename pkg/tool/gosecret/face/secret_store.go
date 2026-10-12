package face

type SecretStore interface {
	Put(
		path string,
		values map[string]string,
	) ([]string, error)
	Reveal(
		path string,
		key string,
	) (string, error)
}
