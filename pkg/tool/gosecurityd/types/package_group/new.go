package package_group

func New(
	source string,
	installed string,
	candidate string,
) *Group {
	return &Group{Source: source, Installed: installed, Candidate: candidate}
}
