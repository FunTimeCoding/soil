package option

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/browser/ladybird"
	"github.com/funtimecoding/soil/pkg/browser/types/ladybird_option"
	"github.com/funtimecoding/soil/pkg/browser/types/message"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
)

func New() *Ladybird {
	result := &Ladybird{Profile: constant.Profile, Headless: true}
	result.Client = func(events func(*message.Reply)) face.LadybirdClient {
		o := ladybird_option.New()
		o.Profile = result.Profile
		o.Headless = result.Headless
		o.Event = events
		c := ladybird.New(o)
		c.Open()

		return c
	}

	return result
}
