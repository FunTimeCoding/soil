package gomattermost

import (
	"context"
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	generativeFace "github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/soil/pkg/text/template"
	"github.com/funtimecoding/soil/pkg/tool/gomattermost/thread_summary"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"time"
)

func Summarize(
	c face.Client,
	o generativeFace.OllamaClient,
	channel string,
	since time.Time,
) ([]*thread_summary.Summary, error) {
	relevant, e := c.GetUnresolvedThreads(
		context.Background(),
		channel,
		&client.GetUnresolvedThreadsParams{Since: since},
	)

	if e != nil {
		return nil, e
	}

	f := chatConstant.MattermostFormat
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
	result := []*thread_summary.Summary{}

	for _, h := range relevant {
		r, g := o.Generate(
			generate_request.New().Prompt(
				template.Execute(tem, formatThread(f, h)),
			),
		)

		if g != nil {
			return nil, g
		}

		result = append(result, thread_summary.New(h, r.Text))
	}

	return result, nil
}
