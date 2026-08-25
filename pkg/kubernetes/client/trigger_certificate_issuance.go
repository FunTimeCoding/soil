package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"time"
)

func (c *Client) TriggerCertificateIssuance(namespace string, name string) {
	resource := c.dynamic.Resource(constant.CertificateGVR).Namespace(namespace)
	o, e := resource.Get(c.context, name, v1.GetOptions{})
	errors.PanicOnError(e)
	existing, _, f := unstructured.NestedSlice(
		o.Object,
		constant.StatusField,
		constant.ConditionsField,
	)
	errors.PanicOnError(f)
	conditions := make([]any, 0, len(existing)+1)

	for _, raw := range existing {
		condition, isMap := raw.(map[string]any)

		if isMap &&
			condition[constant.ConditionFieldType] == constant.ConditionIssuing {
			continue
		}

		conditions = append(conditions, raw)
	}

	conditions = append(
		conditions,
		map[string]any{
			constant.ConditionFieldType:    constant.ConditionIssuing,
			constant.ConditionFieldStatus:  constant.ConditionStatusTrue,
			constant.ConditionFieldReason:  constant.ConditionReasonTriggered,
			constant.ConditionFieldMessage: constant.ConditionMessageTriggered,
			constant.ConditionFieldTransition: time.Now().UTC().Format(
				time.RFC3339,
			),
		},
	)
	errors.PanicOnError(
		unstructured.SetNestedSlice(
			o.Object,
			conditions,
			constant.StatusField,
			constant.ConditionsField,
		),
	)
	_, g := resource.UpdateStatus(c.context, o, v1.UpdateOptions{})
	errors.PanicOnError(g)
}
