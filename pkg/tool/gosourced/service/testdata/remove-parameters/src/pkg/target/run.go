package target

func Run(
	o *Option,
	version string,
) string {
	return Mount(o.Name, version, 1)
}
