package rule_list

import (
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
	"github.com/go-test/deep"
	"log"
)

func (l *List) Add(i *rule.Rule) {
	for _, r := range l.rules {
		if i.Name != r.Name || i.Group != r.Group {
			continue
		}

		if i.RawAlert != nil && r.RawAlert != nil {
			if alertDrifted(i.RawAlert, r.RawAlert) {
				log.Panicf(
					"duplicate RawAlert: name=%s comparison=%+v",
					r.Name,
					deep.Equal(i.RawAlert, r.RawAlert),
				)
			}
		} else if i.RawRecord != nil && r.RawRecord != nil {
			if recordDrifted(i.RawRecord, r.RawRecord) {
				log.Panicf(
					"duplicate RawRecord: name=%s comparison=%+v",
					r.Name,
					deep.Equal(i.RawRecord, r.RawRecord),
				)
			}
		} else {
			log.Panicf("duplicate: name=%s existing=%+v new=%+v", r.Name, i, r)
		}
	}

	l.rules = append(l.rules, i)
}
