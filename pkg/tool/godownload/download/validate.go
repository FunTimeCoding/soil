package download

import "github.com/funtimecoding/soil/pkg/tool/godownload/download/option"

func validate(o *option.Download) {
	if o.Host == "" {
		panic("host required")
	}

	if o.Token == "" {
		panic("token required")
	}

	if o.Owner == "" {
		panic("owner required")
	}

	if o.Repository == "" {
		panic("repository required")
	}

	if o.PackageVersion == "" {
		panic("package version required")
	}

	if o.Package == "" {
		panic("package required")
	}
}
