package face

import "github.com/mattermost/mattermost/server/public/model"

type DialogClient interface {
	MustOpenDialog(v model.OpenDialogRequest)
	MustSubmitDialog(v model.SubmitDialogRequest) *model.SubmitDialogResponse
	OpenDialog(v model.OpenDialogRequest) error
	SubmitDialog(v model.SubmitDialogRequest) (*model.SubmitDialogResponse, error)
}
