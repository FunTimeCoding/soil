package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) UploadFile(
	h *model.Channel,
	name string,
	content []byte,
) (*model.FileInfo, error) {
	result, r, e := c.client.UploadFile(c.context, content, h.Id, name)

	if e != nil {
		return nil, wrapError(r, e)
	}

	if len(result.FileInfos) == 0 {
		return nil, emptyUpload(name)
	}

	return result.FileInfos[0], nil
}
