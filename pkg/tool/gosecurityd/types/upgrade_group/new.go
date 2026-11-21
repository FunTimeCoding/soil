package upgrade_group

func New(
	release string,
	source string,
	kernel bool,
) *Group {
	return &Group{Release: release, Source: source, Kernel: kernel}
}
