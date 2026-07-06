package sed

import "github.com/funtimecoding/soil/pkg/tool/gosed/sed/option"

func validate(o *option.Sed) {
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

	if o.Message == "" {
		panic("message required")
	}

	hasPath := o.Path != ""
	hasActions := len(o.RawActions) > 0

	if hasPath && hasActions {
		panic("cannot mix --path with --action")
	}

	if !hasPath && !hasActions {
		panic("either --path with --replace or --action is required")
	}

	if hasPath && len(o.Replaces) == 0 {
		panic("at least one replace required")
	}
}
