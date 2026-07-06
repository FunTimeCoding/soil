package prompts

import "github.com/funtimecoding/soil/pkg/generative/ollama/decide_action_prompt"

func DecideAction() *decide_action_prompt.Prompt {
	p := decide_action_prompt.New()
	p.Instructions(
		`You are a ChatOps bot.
You will be presented with the current snapshot of messages in a channel.
Your job is to decide whether to take action or not.
Your abilities are:
1. You can give a useful answer to the user, if the question fits into your role as a ChatOps bot.
2. You can create a new issue in the issue tracker, if the user asks you to do so.
If the latest message in the channel does not fit into your role as a ChatOps bot, you should not take action. To indicate this, you should return the string "No action needed".`,
	)

	return p
}
