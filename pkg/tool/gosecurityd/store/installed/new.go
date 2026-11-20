package installed

func New(
	name string,
	version string,
	source string,
	sourceVersion string,
	candidate string,
) *Installed {
	return &Installed{
		Name:          name,
		Version:       version,
		Source:        source,
		SourceVersion: sourceVersion,
		Candidate:     candidate,
	}
}
