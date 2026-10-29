package service_tester

func OverrideManifest() string {
	return "apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: kilo\ndata:\n  key: value\n"
}
