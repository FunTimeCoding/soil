package reaction

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mattermost/mattermost/server/public/model"
	"log"
)

func Decode(v *model.WebSocketEvent) *model.Reaction {
	a, anyOkay := v.GetData()[constant.MattermostReactionField]

	if !anyOkay {
		log.Panicf("no reaction field")
	}

	reaction, castOkay := a.(string)

	if !castOkay {
		log.Panicf("reaction field not string %T", a)
	}

	var result *model.Reaction
	notation.MustDecode(reaction, &result, false)

	return result
}
