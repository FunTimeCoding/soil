package constant

const RemediationGuideLocator = "gosecurityd://guide/remediation"
const RemediationGuide = `# Remediation

The remediation list answers what has to be done so that every
vulnerability at or above a score (9.0, critical, by default) is addressed.

## Addressed

A vulnerability is addressed when it is fixed on every host, or when a
decision records a mitigation or an accepted risk with its reason. The
remediation list is computed on every read; a fixed item leaves it, and a
host that regresses brings it back. Nothing is created or closed behind
anyone's back.

## One item per unit of action

Each finding at or above the score goes into exactly one remediation item,
checked in this order:

| Finding | Item | Subject |
|---|---|---|
| Debian marks it unimportant | none; counted as left out | - |
| the host's release is unsupported or untracked | the host's release upgrade | release:<host> |
| a newer installed kernel has the fix | a reboot | reboot:<host> |
| an upgrade on this host fixes it | the upgrade, shared by all hosts it reaches | upgrade:<release>:<source> |
| a kernel fix Debian has, which nothing in the host's sources reaches | the host's kernel, with the reason (kernel metapackage not installed, no backports source, package lists N days old, or what the sources offer) | kernel:<host> |
| anything else | the vulnerability, with hosts per fix kind | its CVE identifier |

An untracked host is a release upgrade whose vulnerabilities are unknown.
Items list open ones first, then the most hosts times vulnerabilities.

## Decisions

A decision is a flag keyed by an item's subject. flag_vulnerability records
one with a note; the item shows as flagged and stays open. mitigate_flag
closes it with how - "accepted risk: retires in the first quarter, owner
X" or "mitigated: port closed on the load balancer"; the item then counts
as addressed. Flagging a mitigated subject again reopens it. A
vulnerability subject also stores its score when flagged; other subjects
do not consult the database.
`
