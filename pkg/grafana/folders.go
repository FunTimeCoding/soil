package grafana

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/grafana/grafana-openapi-client-go/client/folders"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func (c *Client) Folders() []*models.FolderSearchHit {
	r, e := c.client.Folders.GetFolders(&folders.GetFoldersParams{})
	errors.PanicOnError(e)

	return r.Payload
}
