package commit

import "github.com/funtimecoding/soil/pkg/tool/gocommit/commit/option"

func validate(o *option.Commit) {
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

	if o.Branch == "" {
		panic("branch required")
	}

	if o.Path == "" {
		panic("path required")
	}

	if o.Template == "" {
		panic("template required")
	}

	if o.Message == "" {
		panic("message required")
	}
}
