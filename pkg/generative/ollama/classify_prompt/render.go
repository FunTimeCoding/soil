package classify_prompt

import "github.com/funtimecoding/soil/pkg/text/multi_line"

func (p *Prompt) Render() string {
	m := multi_line.New()
	m.Format("Instructions: %s", p.instructions)
	m.Format("Sample to classify: %s", p.toClassify)
	m.Blank()
	m.Add("## Answer format")
	m.Add(p.answerFormat)

	for _, c := range p.choices {
		m.Add(c)
	}

	m.Blank()
	m.Add("## Example data")

	for _, e := range p.examples {
		m.Add(e)
	}

	return m.Render()
}
