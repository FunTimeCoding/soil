package gomattermost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/argument"
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/text/template"
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermost/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
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
	t := timeLibrary.Midnight(time.Now())
	console.Format(
		"Unresolved threads since %s\n",
		t.Format(timeConstant.DateMinute),
	)
	relevant, e := c.GetUnresolvedThreads(
		context.Background(),
		environment.Required(chatConstant.MattermostChannelEnvironment),
		&generated.GetUnresolvedThreadsParams{Since: t},
	)

	if e != nil {
		n.Exitln(e)

		return
	}

	f := chatConstant.MattermostFormat
	o := ollama.NewEnvironment()
	defer o.Close()
	te := `
Summarize this tech support chat thread.
Identify:
- the MOST RECENT unclear question of the ongoing investigation (paraphrase it)
- the next step SOMEONE needs to do to get closer to a solution, concluded from the conversation people are having (paraphrase it)
Respond in the expected summary format, just the two lines and without extra styling.

## Expected summary format

Latest unclear question: [most recent problem]
Next step: [the concluded next step, or "probably resolved", or "unclear"]

## Example responses from previous threads

### Example 1 (concluded next step)

Latest unclear question: Why is the server not responding?
Next step: jdoe checks the servers logs

### Example 2 (probably resolved)

Latest unclear question: It looks like it works now
Next step: probably resolved

### Example 3 (unclear next step)

Latest unclear question: I don't understand why the database is not connecting
Next step: unclear

## Thread log

{{.}}
`
	tem := template.New("simple", te)

	for _, h := range relevant {
		formatted := formatThread(f, h)
		console.Line(formatted)
		r := o.GenerateSimple(template.Execute(tem, formatted))
		console.Format("%s\n", consoleConstant.Magenta("%s", r.Text))
	}

	if len(relevant) == 0 {
		console.Line("No unresolved threads")
	}
}
