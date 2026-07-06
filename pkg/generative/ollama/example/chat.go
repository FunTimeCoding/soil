package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_request"
)

func Chat() {
	o := ollama.NewEnvironment()
	r := chat_request.New()
	f := option.Color
	r.System(
		"You are a ChatOps bot. Respond in the users message length and writing style.",
	)
	r.User("My server is down.")
	fmt.Println("Messages 1:")
	r.Print()
	fmt.Println()
	a1 := o.MustChat(r)
	fmt.Printf("Response 1: %+v\n", a1.Format(f))
	fmt.Println()
	r.Message(a1.Message)
	r.User("What else can you suggest?")
	fmt.Println("Messages 2:")
	r.Print()
	fmt.Println()
	a2 := o.MustChat(r)
	fmt.Printf("Response 2: %+v\n", a2.Format(f))
}
