package rule_list

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"github.com/go-test/deep"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"log"
)

func (l *List) Add(i *rule.Rule) {
	for _, r := range l.rules {
		if i.Name != r.Name || i.Group != r.Group {
			continue
		}

		if i.RawAlert != nil && r.RawAlert != nil {
			if !LabelsSame(i.RawAlert, r.RawAlert) {
				if false {
					// This is fine
					fmt.Printf(
						"labels differ: %+v %+v\n",
						i.RawAlert.Labels,
						r.RawAlert.Labels,
					)
				}
			} else if i.RawAlert.Query != r.RawAlert.Query {
				if false {
					// This is fine
					fmt.Printf(
						"queries differ: %s %s\n",
						i.RawAlert.Query,
						r.RawAlert.Query,
					)
				}
			} else {
				if !cmp.Equal(
					i.RawAlert,
					r.RawAlert,
					cmpopts.IgnoreFields(
						v1.AlertingRule{},
						"EvaluationTime",
						"LastEvaluation",
					),
				) {
					log.Panicf(
						"duplicate RawAlert: name=%s comparison=%+v",
						r.Name,
						deep.Equal(i.RawAlert, r.RawAlert),
					)
				}
			}
		} else if i.RawRecord != nil && r.RawRecord != nil {
			if c := deep.Equal(i.RawAlert, r.RawAlert); len(c) != 0 {
				log.Panicf(
					"duplicate RawRecord: name=%s comparison=%+v",
					r.Name,
					c,
				)
			}
		} else {
			log.Panicf(
				"duplicate: name=%s existing=%+v new=%+v",
				r.Name,
				i,
				r,
			)
		}
	}

	l.rules = append(l.rules, i)
}
