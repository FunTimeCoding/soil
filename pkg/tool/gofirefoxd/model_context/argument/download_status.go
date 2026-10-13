package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type DownloadStatus struct {
	DownloadIdentifier integer.Integer `json:"download_id"`
}
