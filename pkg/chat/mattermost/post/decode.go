package post

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mattermost/mattermost/server/public/model"
	"log"
)

func Decode(v *model.WebSocketEvent) *model.Post {
	a, anyOkay := v.GetData()[constant.MattermostPostField]

	if !anyOkay {
		log.Panicf("no post field")
	}

	post, castOkay := a.(string)

	if !castOkay {
		log.Panicf("post field not string %T", a)
	}

	var result *model.Post
	notation.MustDecode(post, &result, false)

	return result
}
