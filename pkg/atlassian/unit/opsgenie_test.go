package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/override"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/rotation"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"testing"
)

func TestUserInvolved(t *testing.T) {
	assert.False(
		t,
		opsgenie.UserInvolved(
			stringsConstant.UpperAlfa,
			[]*override.Override{{User: stringsConstant.UpperBravo}},
			[]*rotation.Rotation{},
		),
	)
	assert.True(
		t,
		opsgenie.UserInvolved(
			stringsConstant.UpperAlfa,
			[]*override.Override{},
			[]*rotation.Rotation{
				{
					Participants: []og.Participant{
						{Username: stringsConstant.UpperAlfa},
					},
				},
			},
		),
	)
}
