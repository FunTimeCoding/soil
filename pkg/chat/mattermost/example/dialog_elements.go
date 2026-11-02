package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/mattermost/mattermost/server/public/model"
)

func dialogElements() {
	elements := []model.DialogElement{
		{
			DisplayName: "",
			Name:        "",
			Type:        "",
			SubType:     "",
			Default:     "",
			Placeholder: "",
			HelpText:    "",
			Optional:    false,
			MinLength:   0,
			MaxLength:   0,
			DataSource:  "",
			Options:     nil,
		},
	}
	console.Format("Dialog elements: %d\n", len(elements))
}
