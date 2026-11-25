---
base: pkg/tool/goatlassiand
---

# HTML Web Package

A service tool serving HTML (gomponents) groups its routes under `web/` rather
than `server/`. `web/` holds the `Server` - its services and a view - with its
`New`, `Mount`, `Recovery` and the embedded favicon. The handlers stay methods
on that one `Server`: each fetches through a service and hands the result to a
page package, `<page>_page`, which only renders. Handlers are named after the
route with no `handle` prefix (`dashboard()`, `addSubmit()`); builders after
what they produce (`addForm()`), so the two never collide.

## Routes and guarding

`Mount()` registers every route through the `guard.Mux` verbs with method
values: `g.Open(route.Get("/alerts"), s.alerts)`. Sign-on is in the guard
before `Mount` runs: with `Authorization` set, `o.Web.Guard` installs the
client's session check and opens the auth trio, and protected routes register
with `g.Session`, which puts the signed-in subject on the request
(`guard.Subject`). A service whose session check differs, like gogated's admin
check, calls `g.WithSession` itself. A session route also answers to the
service bearer token - in the `Authorization` header only, never a cookie or a
query string. The `route_guard` analyzer flags a bare `m.HandleFunc` in a web
package.

## Recovery

A web service wires the layout-aware recovery instead of the generic one:
`Run` passes `d.Recovery(r)` to `WithMiddleware`, the daemon delegating to its
web and the web to `view.Recovery`. Panics then render into the page layout -
a notification item with the Sentry event ID, a fragment for HTMX requests -
instead of a plain-text 500.

## Form Errors

Validation errors are not panics. A full-page form bounces back through
`form.Redirect` with the message and submitted values, and the page renders
`form.ErrorText` as a `layout.Alert` above the repopulated form. An
HTMX-swapped form re-renders with the alert inside the swap target. The
sentinel-to-message mapping stays in the service, next to the sentinel.

## View and Layout

`New` builds the view over a `layout.New(identity, theme, favicon)` with the
daemon's style and navigation items; page titles and paths are constants both
the items and the handlers use. Handlers call the view directly -
`RenderPage`, `RenderFragment` behind `IsExtendedRequest` - with no wrappers of
their own. The identity's name links to `/`, so the root takes no navigation
item.

The layout writes the shared component style ahead of the daemon's own. Its
command palette lists the navigation items, `WithCommands` adds the rest, and a
signed session also gets Sign out; a layout with neither, like gogated's
sign-in pages, has no palette. `s.view.Mount(g)` serves the favicon open and
the palette behind sign-on when the guard has a session.

`layout.WithSummary` adds a dot-separated status strip, and
`layout.WithLiveEndpoint` connects `main` to an SSE endpoint, through which the
strip updates (`layout.SummaryStripContent`, event
`constant.LayoutSummaryStrip`). Themes are `Theme*` constants in
`pkg/web/constant`, Pico custom property overrides; downstream repos can add
their own.
