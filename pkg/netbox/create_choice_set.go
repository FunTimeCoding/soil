package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set_record"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set_request"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) CreateChoiceSet(
	name string,
	choices []string,
) (*choice_set.Set, error) {
	q := choice_set_request.New(choices)
	q.Name = name
	result := choice_set_record.New()

	if e := c.rawRequest(
		http.MethodPost,
		constant.ChoiceSetPath,
		q,
		result,
	); e != nil {
		return nil, e
	}

	return choice_set.New(result), nil
}
