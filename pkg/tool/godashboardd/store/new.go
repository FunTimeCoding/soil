package store

func New(
	postgresLocator string,
	litePath string,
) *Store {
	if postgresLocator != "" {
		return newPostgres(postgresLocator)
	}

	if litePath != "" {
		return newLite(litePath)
	}

	panic("set postgresLocator or litePath")
}
