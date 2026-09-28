package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/library"
	"github.com/funtimecoding/soil/pkg/jellyfin/response"
)

func (c *Client) Libraries() ([]*library.Library, error) {
	var out response.Libraries
	e := c.get("/Library/MediaFolders", nil, &out)

	if e != nil {
		return nil, e
	}

	result := make([]*library.Library, len(out.Items))

	for i, r := range out.Items {
		result[i] = library.New(r.Identifier, r.Name, r.CollectionType)
	}

	return result, nil
}
