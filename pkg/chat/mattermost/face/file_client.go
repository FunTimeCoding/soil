package face

import "github.com/mattermost/mattermost/server/public/model"

type FileClient interface {
	File(identifier string) ([]byte, error)
	FileInformation(identifier string) (*model.FileInfo, error)
	UploadFile(
		h *model.Channel,
		name string,
		content []byte,
	) (*model.FileInfo, error)
}
