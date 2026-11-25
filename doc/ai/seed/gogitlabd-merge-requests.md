# gogitlabd merge requests

Recent merge requests for gogitlabd, and watching them from a session
through goclauded. Step 1 is built, with the deviations noted in it;
step 2 is not. An item leaves this seed when it is done.

## Step 1: per-project listing and merge request watches

### Listing

`list_merge_requests` keeps its name and gains:

| Parameter | Default |
|---|---|
| `limit` | 10, at most 100 |
| `offset` | 0 |
| `state` | `opened` (`opened`, `merged`, `closed`, `all`) |
| `order_by` | `updated_at` (or `created_at`), newest first |
| `updated_after` | none; a duration (`7d`, `36h`) or RFC3339 |
| `author`, `reviewer`, `assignee` | none; usernames |
| `source_branch`, `target_branch` | none |
| `search` | none; title and description |
| `draft` | none |

- **Paging:** GitLab pages; the client turns `limit` and `offset` into
  `page`/`per_page`, fetches at most two pages and returns exactly `limit`
  items. Today's fetch-every-page loop goes.
- **Fields:** the answer adds author, reviewers, source and target
  branch, updated time, draft, `detailed_merge_status` and
  `user_notes_count`. The list endpoint already returns them all.

### Watches

Tools, mirroring gomattermostd's `subscribe_thread`:
- `subscribe_merge_request(project, merge_request, callsign, alias,
  delivery)`
- `unsubscribe_merge_request(project, merge_request, callsign)`
- `list_subscriptions(callsign)`

**Store.** gogitlabd's first database: `a.Database()` / `--lite`,
`migrate`, `store`.
- One row per (callsign, instance, project, IID) with alias and delivery.
- One snapshot per merge request, shared by every subscriber:
  - `updated_at`, state and draft;
  - `detailed_merge_status`;
  - head pipeline ID and status;
  - commit `sha`;
  - `user_notes_count` and the highest note ID seen;
  - approved-by.
- A restart reloads the snapshot and announces nothing old.

**Clock.** Every 60 s, per watched merge request:
1. The merge request itself.
2. Its approvals: a new client method, since approvals don't move
   `updated_at`.
3. Its notes newer than the last seen ID, only when the notes count or
   `updated_at` moved.

**Events from the snapshot difference:**
- new comments (author and excerpt);
- approved or unapproved;
- pipeline status changed;
- commits pushed;
- merged or closed. Reopening is never seen, since closing ends the
  watch;
- draft toggled;
- mergeability changed. GitLab passes through transient statuses
  (`checking`, `unchecked`, `preparing`, `approvals_syncing`) while it
  recomputes; the snapshot keeps the last settled one, so only a settled
  change is reported.

**Delivery.**
- One digest per merge request and poll, sent right away; no quiet
  window.
- Through goclauded's `Notify(callsign, "gogitlabd", body, delivery)`, the
  daemon's identity as source like gomattermostd's, behind
  a peer guard, with the pending retry.
- Without goclauded configured, subscribing answers `not_configured`.

**Own comments.**
- `create_merge_request_note` gains an optional `callsign`.
- gogitlabd remembers the note ID it created and the callsign that asked
  for it.
- That note is left out of the writer's digest and shown to other
  subscribers, marked with the writing session.
- Everything else by the token's user (automations, the web UI) is shown.

**End of a watch.**
- **Merged or closed:** the last digest goes out with the subscriber's
  delivery, then a queued line ("!12 merged by …, stopped watching"), and
  the subscription is deleted.
- **Session gone:** a callsign absent from goclauded's roster for 24 h
  loses its watches. An hourly sweep makes one request to goclauded's new
  `GET /api/roster?seen_within=24h`, which answers every callsign seen in
  the window with `last_seen` and `on_roster`. The rule is the same as for
  the roster tool: a session that holds a callsign counts.

### Shared notification code

gomattermostd's `notifier`, `pending_digest` and the general parts of
`digest` move into one shared soil package. gomattermostd and gogitlabd
both use it, and gographd can use it for Teams later. gomattermostd
moves onto it in the same change.

## Step 2: listing across projects

- **Tool:** `list_merge_requests` takes `project`, `group` or neither.
  Neither lists across the instance with `scope=all`. Without a path,
  `instance` picks the GitLab (default instance otherwise).
- **Filter:** a `merge_request_filter` value carries limit, offset,
  state, order and every filter. It turns itself into the project, group
  or instance option struct.
- **"me":** `author`, `reviewer` and `assignee` accept `me`, the token's
  user.
- **Client:** one method `MergeRequests(f)` picks the endpoint and does
  the paging. `AssignedMergeRequests` and `ReviewingMergeRequests` become
  presets of it for the board and leave the face and the mock.
- **Service:** resolves project and group names (a group resolver beside
  the project one).

## Later, not planned

- Watching a whole project: a cheap `updated_after` feed of opened,
  merged and closed merge requests.
- Webhooks: a separate design. Polling stays until then.
