package constant

const KubernetesWorkflowURI = "gokubernetesd://guide/kubernetes-workflow"
const KubernetesWorkflow = `# Kubernetes Workflow

## Cluster selection

Call list_clusters and use_cluster before other operations.
The active cluster is session-scoped.

## Response format

Single resources return YAML. Lists return structured summaries
with restart counts when non-zero. Noisy fields (managedFields,
last-applied-configuration) are stripped by default - a comment
shows what was filtered. Pass unfiltered: true for the raw object.

## Log resolution

name accepts resource/name syntax: deployment/sentry,
job/migrate, cronjob/backup. When multiple pods match, the
error lists them - pick one or pass all: true. When multiple
containers exist, the error names them - pass container.

## Resource usage

top shows CPU and memory for nodes (with percentages) or pods
(raw usage). Pass containers: true for per-container breakdown
on multi-container pods.

## Event muting

Events are filtered against a mute list. Use mute_event to
suppress noisy reasons (e.g. DNSConfigForming). Rules can target
a specific cluster or all clusters. Pass unfiltered: true on
events to see muted events. Mute rules persist across restarts.

## ArgoCD

argocd lists all applications with sync/health status. Pass name
for detail. When synced, bulk fields (resources, history, sync
result resources) are stripped - they appear automatically when
an application is OutOfSync.

## Certificates

certificates lists all certs sorted by expiry with renewal health.
A cert showing "Ready" with "renewal: Failed" means the current
cert works but renewal is stuck. Pass name + namespace for detail
including the latest CertificateRequest.

## Patch and apply

patch modifies an existing resource with a JSON body. Defaults
to strategic merge patch. apply creates a resource from a YAML
manifest - fails if the resource exists unless override: true.
Both support dryRun: true to preview. These are write tools,
hidden by --read-only.
`
