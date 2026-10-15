package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set_record"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) ChoiceSets() ([]*choice_set.Set, error) {
	var result choice_set_record.Page

	if e := c.rawRequest(
		http.MethodGet,
		constant.ChoiceSetListPath,
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return choice_set.NewSlice(result.Results), nil
}
