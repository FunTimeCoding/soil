package face

type Labeller interface {
	ClaimLabel(
		sessionIdentifier string,
		key string,
		value string,
		from string,
	) (bool, error)
	AddLabel(
		sessionIdentifier string,
		key string,
		value string,
		from string,
	) error
	ReleaseLabel(
		sessionIdentifier string,
		key string,
		value string,
		from string,
	) (bool, error)
}
