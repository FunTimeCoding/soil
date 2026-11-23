package selection_filter

import "net/url"

func (f *Filter) Keep(v url.Values) *Filter {
	f.kept = v

	return f
}
