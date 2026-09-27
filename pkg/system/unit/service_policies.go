package unit

import "github.com/funtimecoding/soil/pkg/system/service"

func servicePolicies() map[string]*service.Policy {
	return map[string]*service.Policy{
		"foxtrot": {Origin: "apt.example.org stable/main", Version: "1.2.3"},
	}
}
