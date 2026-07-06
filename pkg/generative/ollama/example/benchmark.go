package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"strings"
	"time"
)

func Benchmark() {
	o := ollama.NewEnvironment()
	base := "Please summarize the following text:\n"
	chunk := "This is a test paragraph with about 50 tokens of content repeated multiple times."

	for i := 1; i <= 5; i++ {
		prompt := fmt.Sprintf("%s%s", base, strings.Repeat(chunk, i*100))
		start := time.Now()
		response := o.GenerateSimple(prompt)
		duration := time.Since(start)
		fmt.Printf(
			"Attempt %d: Response: %s, Duration: %s, Tokens: %.0f\n",
			i,
			response.Text,
			duration,
			response.Tokens,
		)
	}
}
