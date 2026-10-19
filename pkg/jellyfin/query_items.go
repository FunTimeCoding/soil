package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/basic"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/jellyfin/response"
	"net/url"
)

func queryItems(
	b *basic.Client,
	path string,
	p url.Values,
) ([]*item.Item, int, error) {
	var out response.Items
	e := b.Get(path, p, &out)

	if e != nil {
		return nil, 0, e
	}

	result := make([]*item.Item, len(out.Items))

	for i := range out.Items {
		result[i] = toItem(&out.Items[i])
	}

	return result, out.TotalRecordCount, nil
}
