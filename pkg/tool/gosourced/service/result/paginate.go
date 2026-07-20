package result

import "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"

func (r *References) Paginate(limit int, offset int) {
	if limit <= 0 {
		limit = constant.ReferenceLimit
	}

	if offset >= len(r.Locations) {
		r.Locations = nil
		r.More = 0

		return
	}

	r.Locations = r.Locations[offset:]

	if len(r.Locations) > limit {
		r.More = len(r.Locations) - limit
		r.Locations = r.Locations[:limit]
	}
}
