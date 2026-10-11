package service_tester

import "k8s.io/apimachinery/pkg/apis/meta/v1"

func apiResources() []*v1.APIResourceList {
	return []*v1.APIResourceList{
		{
			GroupVersion: "v1",
			APIResources: []v1.APIResource{
				{Name: "pods", Namespaced: true, Kind: "Pod"},
				{
					Name:       "services",
					Namespaced: true,
					Kind:       "Service",
					ShortNames: []string{"svc"},
				},
				{Name: "nodes", Namespaced: false, Kind: "Node"},
				{Name: "events", Namespaced: true, Kind: "Event"},
				{
					Name:       "namespaces",
					Namespaced: false,
					Kind:       "Namespace",
					ShortNames: []string{"ns"},
				},
				{
					Name:       "configmaps",
					Namespaced: true,
					Kind:       "ConfigMap",
					ShortNames: []string{"cm"},
				},
				{Name: "secrets", Namespaced: true, Kind: "Secret"},
			},
		},
		{
			GroupVersion: "apps/v1",
			APIResources: []v1.APIResource{
				{
					Name:       "deployments",
					Namespaced: true,
					Kind:       "Deployment",
					ShortNames: []string{"deploy"},
				},
				{
					Name:       "daemonsets",
					Namespaced: true,
					Kind:       "DaemonSet",
					ShortNames: []string{"ds"},
				},
				{
					Name:       "statefulsets",
					Namespaced: true,
					Kind:       "StatefulSet",
					ShortNames: []string{"sts"},
				},
			},
		},
		{
			GroupVersion: "batch/v1",
			APIResources: []v1.APIResource{
				{Name: "jobs", Namespaced: true, Kind: "Job"},
				{Name: "cronjobs", Namespaced: true, Kind: "CronJob"},
			},
		},
		{
			GroupVersion: "networking.k8s.io/v1",
			APIResources: []v1.APIResource{
				{
					Name:       "networkpolicies",
					Namespaced: true,
					Kind:       "NetworkPolicy",
					ShortNames: []string{"netpol"},
				},
			},
		},
		{
			GroupVersion: "argoproj.io/v1alpha1",
			APIResources: []v1.APIResource{
				{Name: "applications", Namespaced: true, Kind: "Application"},
			},
		},
		{
			GroupVersion: "cert-manager.io/v1",
			APIResources: []v1.APIResource{
				{Name: "certificates", Namespaced: true, Kind: "Certificate"},
				{
					Name:       "certificaterequests",
					Namespaced: true,
					Kind:       "CertificateRequest",
				},
			},
		},
		{
			GroupVersion: "metrics.k8s.io/v1beta1",
			APIResources: []v1.APIResource{
				{Name: "nodes", Namespaced: false, Kind: "NodeMetrics"},
				{Name: "pods", Namespaced: true, Kind: "PodMetrics"},
			},
		},
	}
}
