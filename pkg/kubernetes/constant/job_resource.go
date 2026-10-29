package constant

import "k8s.io/apimachinery/pkg/runtime/schema"

var JobResource = schema.GroupResource{Group: "batch", Resource: "jobs"}
