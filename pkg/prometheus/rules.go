package prometheus

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"reflect"
)

func (c *Client) Rules() (*rule_list.List, error) {
	rules, e := c.client.Rules(c.context, nil)

	if e != nil {
		return nil, e
	}

	result := rule_list.New()

	for _, group := range rules.Groups {
		for _, r := range group.Rules {
			switch t := r.(type) {
			case v1.AlertingRule:
				result.Add(rule.NewAlert(&t, &group))
			case v1.RecordingRule:
				result.Add(rule.NewRecord(&t, &group))
			default:
				fmt.Printf("  Unexpected: %s", reflect.TypeOf(r))
			}
		}
	}

	return result, nil
}
