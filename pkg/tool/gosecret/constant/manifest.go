package constant

const ExternalSecretFormat = `---
apiVersion: external-secrets.io/v1
kind: ExternalSecret
metadata: {name: %[1]s}
spec:
  refreshInterval: 1h
  secretStoreRef: {kind: ClusterSecretStore, name: %[3]s}
  target: {name: %[1]s, creationPolicy: Owner%[4]s}
  dataFrom:
  - extract: {key: %[2]s}
`

const TemplateTypeFormat = ", template: {type: %s}"
