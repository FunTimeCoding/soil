package generate_response

import "github.com/funtimecoding/soil/pkg/console"

func (r *Response) Print() {
	console.Format("Total: %dms\n", r.Total)
	console.Format("  Load: %dms\n", r.Load)
	console.Format("  Prompt evaluation: %dms\n", r.PromptEvaluation)
	console.Format("    Tokens/s: %.0f\n", r.PromptTokens)
	console.Format("  Evaluation: %dms\n", r.Evaluation)
	console.Format("    Tokens/s: %.0f\n", r.Tokens)
}
