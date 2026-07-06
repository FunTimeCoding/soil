package timeline

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func FormatDescription(e *Entry) string {
	actor := e.Actor

	if e.Alias != "" {
		actor = fmt.Sprintf("%s (%s)", e.Actor, e.Alias)
	}

	switch e.Kind {
	case "announce":
		return fmt.Sprintf("%s announced: %s", actor, e.meta(constant.Topic))
	case "complete":
		if m := e.meta(constant.Message); m != "" {
			return fmt.Sprintf(
				"%s completed %s: %s",
				actor,
				e.meta(constant.Topic),
				m,
			)
		}

		return fmt.Sprintf("%s completed %s", actor, e.meta(constant.Topic))
	case "update":
		if m := e.meta(constant.Message); m != "" {
			return fmt.Sprintf(
				"%s updated %s: %s",
				actor,
				e.meta(constant.Topic),
				m,
			)
		}

		return fmt.Sprintf(
			"%s updated scope: %s",
			actor,
			e.meta(constant.Topic),
		)
	case "send":
		to := "all"

		if recipient := e.meta("recipient"); recipient != "" {
			to = recipient
		}

		return fmt.Sprintf(
			"%s → %s: %s",
			actor,
			to,
			e.meta(constant.Message),
		)
	case "moment":
		return fmt.Sprintf(
			"%s captured: %s",
			actor,
			e.meta(constant.Line),
		)
	case "summarize":
		body := e.meta(constant.Body)

		if body == "" {
			return fmt.Sprintf("%s summarized", actor)
		}

		if !e.Full && len(body) > 200 {
			body = fmt.Sprintf("%s…", body[:200])
		}

		return fmt.Sprintf("%s summarized: %s", actor, body)
	case "label":
		return fmt.Sprintf("%s label: %s", actor, e.meta("change"))
	case "release":
		return fmt.Sprintf("%s left", actor)
	case "register":
		return fmt.Sprintf("%s joined", actor)
	case "inactivity_timeout":
		return fmt.Sprintf("%s removed from roster (inactive)", actor)
	case "complete_timeout":
		return fmt.Sprintf(
			"%s removed from roster (idle after completing)",
			actor,
		)
	default:
		if name := e.meta(constant.SessionName); name != "" {
			if actor != "" {
				return fmt.Sprintf("%s %s: %s", actor, e.Kind, name)
			}

			return fmt.Sprintf("%s: %s", e.Kind, name)
		}

		if actor != "" {
			return fmt.Sprintf("%s %s", actor, e.Kind)
		}

		return e.Kind
	}
}
