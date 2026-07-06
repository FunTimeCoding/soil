package decide_action_prompt

import "github.com/funtimecoding/soil/pkg/text/multi_line"

func (p *Prompt) Render() string {
	m := multi_line.New()
	m.Format("Instructions: %s", p.instructions)
	m.Blank()
	m.Format("## Messages\n\n%s", p.history)

	return m.Render()
}
