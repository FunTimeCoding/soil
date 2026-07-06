package resource

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/format"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func BuildRequestStatus(
	requests *unstructured.UnstructuredList,
) map[string]string {
	result := make(map[string]string)

	for _, e := range requests.Items {
		owners := ExtractOwnerCertificate(e.Object)

		if owners == "" {
			continue
		}

		key := fmt.Sprintf("%s/%s", e.GetNamespace(), owners)
		ready := ExtractConditionStatus(e.Object)
		age := format.Age(e.GetCreationTimestamp().Time)
		result[key] = fmt.Sprintf("%s (%s)", ready, age)
	}

	return result
}
