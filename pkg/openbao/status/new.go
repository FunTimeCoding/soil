package status

func New(
	initialized bool,
	sealed bool,
	version string,
) *Status {
	return &Status{Initialized: initialized, Sealed: sealed, Version: version}
}
