package constant

const MatchingGuideLocator = "gosecurityd://guide/matching"
const MatchingGuide = `# Matching

How gosecurityd turns a system's installed packages into findings.

## Sources

- **National Vulnerability Database**: a vulnerability's description,
  references and score. gosecurityd mirrors one score per vulnerability
  (scoring version 4.0, else 3.1, 3.0, 2; the database's own score before
  the reporting organisation's) and keeps it current hourly;
  get_score_status shows how far the mirror has come. A vulnerability the
  database has not analysed yet has no score and ranks last.
- **Debian's security tracker** for bookworm, trixie, forky and sid.
- **Freexian's extended tracker** for stretch, buster and bullseye. Its
  fixes sit in paid suites, so a fix there reads "paid support only"
  unless the host's own upgrade candidate already has it.

Debian backports fixes without changing the upstream version, so judge a
Debian host by get_debian_status, never by upstream version ranges.
Debian's urgency is mostly "not yet assigned": rank by score instead.

## Coverage and clean

Every system has one coverage:

- **supported**: a Debian-tracked release; a fix is an upgrade away
- **unsupported, tracked**: a Freexian-tracked release; fixes mostly paid
- **unsupported, untracked**: no tracker (older Debian releases, other
  distributions); every
  package is unchecked and its vulnerabilities are unknown, never zero

A system is **clean** only when it is supported, its tracker is loaded and
verified recently, its snapshot is fresh (stale marks one past its age) and
it has no finding. Under any other coverage, an empty finding list means
unchecked.

## Findings and fix kinds

A finding is one vulnerability on one source package of one system. Each
has one kind:

- **reboot fixes it**: kernel only; a newer installed kernel has the fix
- **upgrade fixes it**: the host's upgrade candidate reaches the fix,
  through any binary package of the source
- **upgrade and reboot fix it**: the same for the kernel
- **Debian fix not in the candidate**: Debian has fixed it, but nothing the
  host's sources offer reaches the fix yet
- **paid support only**
- **no fix yet**

Per vulnerability, a system counts with its worst kind. The candidate is
only as fresh as the host's package lists. Debian marks some
vulnerabilities "unimportant" (it will not fix them); they are marked and
stay out of the remediation list.

## The kernel

Only the running kernel image (uname) is matched, its binary version
against Debian's linux source; headers and other side packages are left
out. A backports kernel is matched against the next release's status. The
kernel has its own state: up to date, reboot fixes it, upgrade and reboot,
no upgrade path (with the reason: kernel metapackage not installed, no
backports source, package lists N days old, or what the sources offer),
paid support only, or no kernel data.
Package counts leave the kernel out.

## The score filter

score is the lowest score to list. It narrows list_systems (systems
carrying a vulnerability at that score), list_upgrades (upgrades fixing
one, counting only those) and list_findings. Critical means 9.0 or higher.
`
