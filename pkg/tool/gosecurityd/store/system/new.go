package system

func New(
	name string,
	release string,
	priority int,
) *System {
	return &System{Name: name, Release: release, Priority: priority}
}
