package permission

import "time"

func NewPrompt(
	session string,
	when time.Time,
	command string,
	reason string,
) *Prompt {
	return &Prompt{
		Session: session,
		When:    when,
		Command: command,
		Reason:  reason,
	}
}
