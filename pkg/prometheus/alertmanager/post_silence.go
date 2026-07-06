package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/system"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func (c *Client) PostSilence(
	identifier string,
	alert string,
	comment string,
	start time.Time,
	end time.Time,
) (string, error) {
	if comment == "" {
		comment = constant.NoComment
	}

	p := silence.NewPostSilencesParams()
	p.Silence = &models.PostableSilence{
		ID: identifier,
		Silence: models.Silence{
			Comment:   &comment,
			CreatedBy: new(system.User().Username),
			Matchers: []*models.Matcher{
				{
					Name:    new(constant.AlertnameLabel),
					Value:   &alert,
					IsRegex: new(false),
				},
			},
			StartsAt: new(library.Scan(start)),
			EndsAt:   new(library.Scan(end)),
		},
	}
	r, e := c.client.Silence.PostSilences(p)

	if e != nil {
		return "", e
	}

	return r.Payload.SilenceID, nil
}
