package face

type Secrets interface {
	StatusSource
	List(path string) ([]string, error)
	Keys(path string) ([]string, error)
	Merge(
		path string,
		values map[string]string,
	) ([]string, error)
	Reveal(
		path string,
		key string,
	) (string, error)
}
