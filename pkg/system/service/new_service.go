package service

func NewService(
	name string,
	state string,
	origin string,
	source string,
	version string,
	deliberate bool,
) *Service {
	return &Service{
		Name:       name,
		State:      state,
		Origin:     origin,
		Source:     source,
		Version:    version,
		Deliberate: deliberate,
	}
}
