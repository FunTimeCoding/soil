package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/override"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/rotation"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"testing"
)

func TestUserInvolved(t *testing.T) {
	assert.False(
		t,
		opsgenie.UserInvolved(
			upper.Alfa,
			[]*override.Override{{User: upper.Bravo}},
			[]*rotation.Rotation{},
		),
	)
	assert.True(
		t,
		opsgenie.UserInvolved(
			upper.Alfa,
			[]*override.Override{},
			[]*rotation.Rotation{
				{
					Participants: []og.Participant{{Username: upper.Alfa}},
				},
			},
		),
	)
}
