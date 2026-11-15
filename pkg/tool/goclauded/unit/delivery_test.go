package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/delivery_tester"
	"slices"
	"strings"
	"testing"
)

func TestDeliveryOrdersSections(t *testing.T) {
	m := delivery_tester.StoredMessage(7, "Dale", "hello")
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
		delivery_tester.Render(
			[]*queue.Entry{
				delivery_tester.MessageEntry(m),
				delivery_tester.Entry(
					constant.QueueMemoryUpdate,
					"thread updated by Glen",
				),
				delivery_tester.Entry(
					constant.QueueNotification,
					"mattermost: a digest",
				),
				delivery_tester.Entry(
					constant.QueueSessionUpdate,
					"Frost updated scope: lint",
				),
				delivery_tester.Entry(
					constant.QueueTimeout,
					"1 hour since last turn.",
				),
				delivery_tester.Entry(constant.QueuePulse, "deploy finished"),
				delivery_tester.Entry(
					constant.QueueReannounce,
					"Re-announce required: MCP binding lost during service restart.",
				),
			},
			m,
		),
	)
}

func TestRenderCutsLongMessage(t *testing.T) {
	m := delivery_tester.StoredMessage(12, "Ash", strings.Repeat("a", 5000))
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
		delivery_tester.Render(
			[]*queue.Entry{delivery_tester.MessageEntry(m)},
			m,
		),
	)
}

func TestRenderWaitsBelowCutFloor(t *testing.T) {
	var entries []*queue.Entry
	stored := delivery_tester.StoredMessages(
		4,
		"Blair",
		strings.Repeat("c", 3000),
	)

	for _, m := range stored {
		entries = append(entries, delivery_tester.MessageEntry(m))
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
		delivery_tester.Render(entries, stored...),
	)
}

func TestDeliveryTrimsMemoryBeforeMessages(t *testing.T) {
	m := delivery_tester.StoredMessage(9, "Dale", "after the memories")
	var entries []*queue.Entry
	var kept []string

	for i := range 70 {
		line := fmt.Sprintf("memory number %02d updated by Ellis", i)
		entries = append(
			entries,
			delivery_tester.Entry(constant.QueueMemoryUpdate, line),
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
		delivery_tester.Render(
			append(entries, delivery_tester.MessageEntry(m)),
			m,
		),
	)
}

func TestDeliveryKeepsMemoryHeaderWhenNothingFits(t *testing.T) {
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
		delivery_tester.Render(
			[]*queue.Entry{
				delivery_tester.Entry(
					constant.QueueMemoryUpdate,
					"thread updated by Glen",
				),
				delivery_tester.Entry(
					constant.QueueMessage,
					join.Empty("Ash: ", body),
				),
			},
		),
	)
}

func TestDeliveryKeepsMessageWithoutRowInline(t *testing.T) {
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
		delivery_tester.Render(
			[]*queue.Entry{
				delivery_tester.Entry(
					constant.QueueMessage,
					join.Empty("Ash: ", body),
				),
			},
		),
	)
}
