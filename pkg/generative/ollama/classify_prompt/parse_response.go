package classify_prompt

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (p *Prompt) ParseResponse(s string) string {
	var result *Response
	notation.MustDecode(s, &result, true)
	result.Classified = p.toClassify

	return fmt.Sprintf("%s\n", notation.Encode(result, true))
}
