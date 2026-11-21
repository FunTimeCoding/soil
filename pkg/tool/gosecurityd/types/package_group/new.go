package package_group

func New(
	source string,
	kind string,
	installed string,
	candidate string,
) *Group {
	return &Group{
		Source:    source,
		Kind:      kind,
		Installed: installed,
		Candidate: candidate,
	}
}
