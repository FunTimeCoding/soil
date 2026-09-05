package service_tester

import (
	"fmt"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

func (t *Tester) AddEvent(
	namespace string,
	reason string,
	message string,
	eventType string,
	involvedKind string,
	involvedName string,
) {
	now := v1.NewTime(time.Now())
	_, e := t.Clientset.CoreV1().Events(namespace).Create(
		t.context(),
		&core.Event{
			ObjectMeta: v1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s", reason, involvedName),
				Namespace: namespace,
			},
			Reason:  reason,
			Message: message,
			Type:    eventType,
			InvolvedObject: core.ObjectReference{
				Kind: involvedKind,
				Name: involvedName,
			},
			LastTimestamp: now,
			Count:         1,
		},
		v1.CreateOptions{},
	)

	if e != nil {
		panic(e)
	}
}
