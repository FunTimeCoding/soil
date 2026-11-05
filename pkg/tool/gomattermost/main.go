package gomattermost

import (
	"github.com/funtimecoding/soil/pkg/argument"
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomattermost/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/client"
	"time"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse()
	n := terminal.New(s)
	c := client.NewEnvironment()
	defer c.Close()
	o := ollama.NewEnvironment()
	defer o.Close()
	since := soilTime.Midnight(time.Now())
	summaries, e := Summarize(
		c,
		o,
		environment.Required(chatConstant.MattermostChannelEnvironment),
		since,
	)

	if e != nil {
		n.EmitDone(e)

		return
	}

	Print(n, since, summaries)
}
