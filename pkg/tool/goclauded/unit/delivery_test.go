package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"slices"
	"strings"
	"testing"
)

func TestDeliveryOrdersTheSections(t *testing.T) {
	m := storedMessage(7, "Dale", "hello")
	assert.String(
		t,
		`Re-announce required: MCP binding lost during service restart.
[pulse] deploy finished
Idle: 1 hour since last turn.
Session activity:
  Frost updated scope: lint
Notifications:
  mattermost: a digest
Memory activity:
  thread updated by Glen
Messages:
  Dale: hello`,
		render(
			[]*queue.Entry{
				messageEntry(m),
				deliveryEntry(
					constant.QueueMemoryUpdate,
					"thread updated by Glen",
				),
				deliveryEntry(
					constant.QueueNotification,
					"mattermost: a digest",
				),
				deliveryEntry(
					constant.QueueSessionUpdate,
					"Frost updated scope: lint",
				),
				deliveryEntry(constant.QueueTimeout, "1 hour since last turn."),
				deliveryEntry(constant.QueuePulse, "deploy finished"),
				deliveryEntry(
					constant.QueueReannounce,
					"Re-announce required: MCP binding lost during service restart.",
				),
			},
			m,
		),
	)
}

func TestDeliveryCutsALongMessage(t *testing.T) {
	m := storedMessage(12, "Ash", strings.Repeat("a", 5000))
	assert.String(
		t,
		join.NewLine(
			[]string{
				"Messages:",
				"  This delivery holds about 2,000 characters; longer messages are cut or waiting, and read_message returns them whole.",
				join.Empty(
					"  Ash: ",
					strings.Repeat("a", 1789),
					"…cut here - 3,211 more characters in message 12, read it with read_message",
				),
			},
		),
		render([]*queue.Entry{messageEntry(m)}, m),
	)
}

func TestDeliveryWaitsBelowTheCutFloor(t *testing.T) {
	var entries []*queue.Entry
	stored := storedMessages(4, "Blair", strings.Repeat("c", 3000))

	for _, m := range stored {
		entries = append(entries, messageEntry(m))
	}

	assert.String(
		t,
		join.NewLine(
			[]string{
				"Messages:",
				"  This delivery holds about 2,000 characters; longer messages are cut or waiting, and read_message returns them whole.",
				join.Empty(
					"  Blair: ",
					strings.Repeat("c", 1518),
					"…cut here - 1,482 more characters in message 1, read it with read_message",
				),
				"  Message 2 from Blair at 14:35 (3,000 characters) is waiting - read it with read_message",
				"  Message 3 from Blair at 14:35 (3,000 characters) is waiting - read it with read_message",
				"  Message 4 from Blair at 14:35 (3,000 characters) is waiting - read it with read_message",
			},
		),
		render(entries, stored...),
	)
}

func TestDeliveryTrimsMemoryBeforeMessages(t *testing.T) {
	m := storedMessage(9, "Dale", "after the memories")
	var entries []*queue.Entry
	var kept []string

	for i := range 70 {
		line := fmt.Sprintf("memory number %02d updated by Ellis", i)
		entries = append(
			entries,
			deliveryEntry(constant.QueueMemoryUpdate, line),
		)

		if i < 52 {
			kept = append(kept, join.Empty("  ", line))
		}
	}

	assert.String(
		t,
		join.NewLine(
			slices.Concat(
				[]string{"Memory activity:"},
				kept,
				[]string{
					"  Memory changes not shown: 18 - list_memories lists recent updates",
					"Messages:",
					"  Dale: after the memories",
				},
			),
		),
		render(append(entries, messageEntry(m)), m),
	)
}

func TestDeliveryKeepsTheMemoryHeaderWhenNothingFits(t *testing.T) {
	body := strings.Repeat("e", 2500)
	assert.String(
		t,
		join.NewLine(
			[]string{
				"Memory activity:",
				"  Memory changes not shown: 1 - list_memories lists recent updates",
				"Messages:",
				"  This delivery holds about 2,000 characters; longer messages are cut or waiting, and read_message returns them whole.",
				join.Empty("  Ash: ", body),
			},
		),
		render(
			[]*queue.Entry{
				deliveryEntry(
					constant.QueueMemoryUpdate,
					"thread updated by Glen",
				),
				deliveryEntry(constant.QueueMessage, join.Empty("Ash: ", body)),
			},
		),
	)
}

func TestDeliveryKeepsAMessageWithoutARowInline(t *testing.T) {
	body := strings.Repeat("d", 2500)
	assert.String(
		t,
		join.NewLine(
			[]string{
				"Messages:",
				"  This delivery holds about 2,000 characters; longer messages are cut or waiting, and read_message returns them whole.",
				join.Empty("  Ash: ", body),
			},
		),
		render(
			[]*queue.Entry{
				deliveryEntry(constant.QueueMessage, join.Empty("Ash: ", body)),
			},
		),
	)
}
