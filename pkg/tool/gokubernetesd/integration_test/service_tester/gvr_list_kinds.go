package service_tester

import "k8s.io/apimachinery/pkg/runtime/schema"

func gvrListKinds() map[schema.GroupVersionResource]string {
	return map[schema.GroupVersionResource]string{
		{Group: "", Version: "v1", Resource: "pods"}:             "PodList",
		{Group: "", Version: "v1", Resource: "services"}:         "ServiceList",
		{Group: "", Version: "v1", Resource: "nodes"}:            "NodeList",
		{Group: "", Version: "v1", Resource: "events"}:           "EventList",
		{Group: "", Version: "v1", Resource: "namespaces"}:       "NamespaceList",
		{Group: "", Version: "v1", Resource: "configmaps"}:       "ConfigMapList",
		{Group: "", Version: "v1", Resource: "secrets"}:          "SecretList",
		{Group: "apps", Version: "v1", Resource: "deployments"}:  "DeploymentList",
		{Group: "apps", Version: "v1", Resource: "daemonsets"}:   "DaemonSetList",
		{Group: "apps", Version: "v1", Resource: "statefulsets"}: "StatefulSetList",
		{Group: "batch", Version: "v1", Resource: "jobs"}:        "JobList",
		{Group: "batch", Version: "v1", Resource: "cronjobs"}:    "CronJobList",
		{
			Group:    "networking.k8s.io",
			Version:  "v1",
			Resource: "networkpolicies",
		}: "NetworkPolicyList",
		{
			Group:    "argoproj.io",
			Version:  "v1alpha1",
			Resource: "applications",
		}: "ApplicationList",
		{Group: "cert-manager.io", Version: "v1", Resource: "certificates"}: "CertificateList",
		{
			Group:    "cert-manager.io",
			Version:  "v1",
			Resource: "certificaterequests",
		}: "CertificateRequestList",
		{Group: "metrics.k8s.io", Version: "v1beta1", Resource: "nodes"}: "NodeMetricsList",
		{Group: "metrics.k8s.io", Version: "v1beta1", Resource: "pods"}:  "PodMetricsList",
	}
}
