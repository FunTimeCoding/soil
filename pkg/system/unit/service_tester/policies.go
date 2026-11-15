package service_tester

import "github.com/funtimecoding/soil/pkg/system/types/service_policy"

func Policies() map[string]*service_policy.Policy {
	return map[string]*service_policy.Policy{
		"foxtrot": {Origin: "apt.example.org stable/main", Version: "1.2.3"},
	}
}
