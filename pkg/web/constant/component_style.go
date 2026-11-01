package constant

const ComponentStyle = `
:where(:root) { --okay: var(--pico-ins-color); --warn: #e8c547; --bad: var(--pico-del-color); --conflict: #f012be; --quiet: var(--pico-muted-color); --busy: #4c9aff; }
:where(form).action { display: inline-block; margin: 0; }
:where(form.action) [type=submit] { width: auto; margin: 0; }
:where(form.action.small) [type=submit] { padding: 0.2rem 0.5rem; font-size: 0.8rem; }
:where(table) .numeric { text-align: right; font-variant-numeric: tabular-nums; }
.absent { color: var(--pico-muted-color); font-style: italic; }
:where(table).facts { max-width: 36rem; }
:where(table).facts th { text-align: left; font-weight: 400; color: var(--pico-muted-color); width: 14rem; }
:where(table).facts tr.section th { font-weight: 600; color: var(--pico-color); padding-top: 0.9rem; }
:where(.tile-grid) { display: grid; grid-template-columns: repeat(auto-fill, minmax(17rem, 1fr)); gap: 1rem; margin: 1rem 0 1.5rem; }
:where(.tile) { border: 1px solid var(--pico-muted-border-color); border-radius: var(--pico-border-radius); padding: 1.25rem; transition: border-color 0.2s; background: var(--pico-card-background-color); }
:where(.tile:hover) { border-color: var(--pico-primary); }
:where(.dial) { position: relative; width: 15rem; height: 15rem; border-radius: 50%; display: flex; align-items: center; justify-content: center; background: var(--dial-background); }
:where(.dial)::after { content: ""; position: absolute; inset: 1.6rem; border-radius: 50%; background: var(--pico-card-background-color); }
:where(.dial-center) { position: relative; z-index: 1; text-align: center; line-height: 1.3; }
:where(.dial) .dial-center strong { display: block; font-size: 1.4rem; }
:where(.dial) .dial-center small { color: var(--pico-muted-color); }
:where(.dial-needle) { position: absolute; top: 50%; left: 50%; width: 50%; height: 0; transform-origin: 0 50%; transform: rotate(var(--dial-angle)); z-index: 1; }
:where(.dial-point) { position: absolute; right: 0.35rem; top: -0.45rem; width: 0.9rem; height: 0.9rem; border-radius: 50%; background: var(--dial-color); box-shadow: 0 0 0.6rem var(--dial-color); }
:where(.dial-tick) { position: absolute; right: 0.55rem; top: -0.25rem; width: 0.5rem; height: 0.5rem; border-radius: 50%; background: var(--dial-color); opacity: 0.8; }
:where(.dial [title]) { cursor: help; }
:where(.status-dot) { display: inline-block; width: 0.65rem; height: 0.65rem; border-radius: 50%; vertical-align: middle; background: var(--dot-color, var(--pico-muted-color)); }
:where(form).entry.inline { display: flex; flex-wrap: wrap; gap: 0.5rem; align-items: end; }
:where(form).entry.inline label { margin: 0; }
:where(form).entry.inline :is(input:not([type=checkbox]), select, textarea) { margin: 0; width: auto; }
:where(form).entry.inline [type=submit] { width: auto; margin: 0; padding: 0.35rem 0.7rem; font-size: 0.85rem; }
:where(form).entry.inline .entry-label { display: block; font-size: 0.8rem; color: var(--pico-muted-color); }
:where(form).entry.record label:has(> .entry-label) { display: grid; grid-template-columns: 14rem 1fr; align-items: center; gap: 0.5rem; margin: 0; padding: 0.6rem 0.5rem; border-bottom: 1px solid var(--pico-muted-border-color); }
:where(form).entry.record .entry-label { font-weight: 600; font-size: 0.95rem; }
:where(form).entry.record :is(input:not([type=checkbox]), select, textarea) { margin: 0; }
:where(form).entry .entry-buttons { display: flex; gap: 0.5rem; margin-top: 1rem; }
:where(form).entry .entry-buttons :is([type=submit], [role=button]) { width: auto; margin: 0; }
:where(form).entry .entry-buttons .entry-alternative { order: -1; margin-right: auto; }
:where(form).entry.inline .entry-buttons [role=button] { padding: 0.35rem 0.7rem; font-size: 0.85rem; }
:where(form).entry.inline .form-alert { flex-basis: 100%; margin: 0; }
:where(form).entry.record.inline { align-items: center; }
:where(form).entry.record.inline label:has(> .entry-label) { flex: 1; padding: 0; border: none; }
:where(form).entry.record.inline :is(input:not([type=checkbox]), select, textarea) { width: 100%; }
:where(form).entry.record.inline .entry-buttons { margin-top: 0; }
:where(.field-row) { display: flex; align-items: center; justify-content: space-between; gap: 0.5rem; padding: 0.6rem 0.5rem; border-bottom: 1px solid var(--pico-muted-border-color); border-radius: var(--pico-border-radius); }
:where(.field-row:last-child) { border-bottom: none; }
:where(.field-row.field-unset) { border-left: 3px solid var(--pico-muted-border-color); background: color-mix(in srgb, var(--pico-background-color) 95%, var(--pico-muted-color)); }
:where(.field-label) { font-weight: 600; min-width: 14rem; font-size: 0.95rem; }
.field-label[data-tooltip] { cursor: help; border-bottom: none !important; text-decoration: none !important; }
:where(.field-value) { flex: 1; font-size: 0.95rem; }
:where(.field-actions) { display: flex; gap: 0.4rem; }
:where(.cluster) { display: flex; flex-wrap: wrap; align-items: center; gap: 0.4rem; }
.field-actions :is(button, [role=button]), .card-task .card-footer :is([type=submit], [role=button]) { padding: 0.35rem 0.7rem; font-size: 0.85rem; margin: 0; }
.field-note { color: var(--pico-secondary); font-style: italic; }
:where(.badge) { --badge-color: var(--pico-primary); display: inline-block; padding: 0.15rem 0.5rem; font-size: 0.75rem; font-weight: 600; border-radius: 1rem; white-space: nowrap; background: color-mix(in srgb, var(--badge-color) 20%, var(--pico-background-color)); color: var(--badge-color); }
:where(.badge.quiet) { --badge-color: var(--pico-muted-color); font-weight: 400; }
.chip { display: inline-block; font-size: 0.75rem; line-height: 1.4; border: 1px solid var(--chip-color, var(--pico-muted-border-color)); border-radius: 0.5rem; padding: 0 0.4rem; margin: 0 0.15rem 0.15rem 0; color: var(--chip-color, inherit); text-decoration: none; }
.chip.toggle { --chip-color: var(--pico-secondary); width: auto; background: transparent; padding: 0.15rem 0.5rem; cursor: pointer; }
.chip.toggle:hover { opacity: 0.8; }
.chip.active { background: var(--pico-primary); border-color: var(--pico-primary); color: var(--pico-primary-inverse); }
:where(.meter) { display: flex; align-items: center; gap: 0.5rem; }
:where(.meter-row) { display: flex; flex-wrap: wrap; align-items: center; gap: 0.6rem; margin-top: 0.5rem; }
:where(.meter-label) { flex: 0 1 7rem; min-width: 0; font-size: 0.875rem; }
:where(.meter-value) { flex: none; min-width: 3rem; font-size: 0.875rem; text-align: right; font-variant-numeric: tabular-nums; }
:where(.meter-icon) { display: inline-flex; align-items: center; opacity: 0.6; }
.meter[data-tooltip] { cursor: help; border-bottom: none !important; text-decoration: none !important; }
progress.meter-bar { width: var(--meter-width, 100%); height: var(--meter-height, 0.5rem); }
.meter progress.meter-bar { margin: 0; flex: 0 1 var(--meter-width, 100%); min-width: min(4rem, var(--meter-width, 4rem)); }
:where(.meter-gauge) { flex: none; height: var(--meter-height, 18px); width: var(--meter-width, 4rem); border-radius: 9px; background: var(--meter-track, var(--pico-muted-border-color)); overflow: hidden; }
:where(.meter-gauge-fill) { height: 100%; width: var(--meter-share); border-radius: 9px; background: var(--meter-color, var(--pico-primary)); transition: width 0.3s; }
.card { position: relative; display: grid; grid-template-columns: auto 1fr; column-gap: 0.9rem; align-content: start; align-items: start; margin: 0; padding: 1rem 1.1rem; border: 1px solid var(--pico-muted-border-color); border-top: 3px solid var(--card-accent, var(--pico-primary)); border-radius: var(--pico-border-radius); transition: border-color 0.15s, transform 0.15s; }
.card:not(:has(> .card-visual)), .card-art, .card-full { grid-template-columns: 1fr; }
.card-body { display: flex; flex-direction: column; min-width: 0; }
.card:hover, .card:focus-within { border-color: var(--card-accent, var(--pico-primary)); transform: translateY(-2px); }
.card-visual { width: 3rem; height: 3rem; border-radius: 50%; overflow: hidden; display: grid; place-items: center; font-size: 1.6rem; background: color-mix(in srgb, var(--card-accent, var(--pico-primary)) 25%, var(--pico-background-color)); }
.card-visual img { display: block; width: 100%; height: 100%; object-fit: cover; object-position: 50% 25%; border-radius: inherit; transition: scale 0.25s; }
.card:hover .card-visual img { scale: 1.04; }
.card-head { display: flex; flex-wrap: wrap; align-items: center; gap: 0.4rem; margin: 0; padding: 0; background: none; border: none; }
.card-head :is(h1, h4) { margin: 0; flex: 1; }
.card-head > :not(h1, h4) { flex: none; }
.card-hero .card-head h1 { font-size: 2rem; }
.card-link, .card-link:hover { color: inherit; text-decoration: none; }
.card-link::after { content: ""; position: absolute; inset: 0; }
.card-subtitle { margin: 0.1rem 0 0; font-size: 0.85rem; }
.card-meta { margin: 0 0 0.3rem; font-size: 0.8rem; color: var(--pico-muted-color); }
.card .badge-row { margin: 0.4rem 0 0; }
.card-footer { position: relative; z-index: 1; display: flex; gap: 0.4rem; margin: 0.6rem 0 0; padding: 0; background: none; border: none; opacity: 0.55; transition: opacity 0.15s; }
.card:hover .card-footer, .card:focus-within .card-footer { opacity: 1; }
.card-footer [role=button] { width: auto; margin: 0; padding: 0.2rem 0.5rem; font-size: 0.8rem; }
.card-task .card-footer { flex-wrap: wrap; opacity: 1; }
.card-compact { padding: 0.6rem 0.8rem; }
.card-compact .card-visual { width: 2rem; height: 2rem; font-size: 1.1rem; }
.card-hero { padding: 1.4rem 1.6rem; }
.card-hero .card-visual { width: 5rem; height: 5rem; font-size: 2.8rem; }
.tile-grid > .card-hero { grid-column: 1 / -1; }
.card-art .card-visual { width: auto; height: auto; aspect-ratio: 4 / 3; margin: -1rem -1.1rem 0.8rem; border-radius: 0; font-size: 7rem; }
.card-side { display: flow-root; }
.card-side .card-body { display: block; }
.card-side .card-visual { float: right; margin: 0 0 0.6rem 0.9rem; width: 7rem; height: 7rem; border-radius: 0.6rem; font-size: 3.5rem; }
.card-side .card-visual img { transform: scale(1.2); transform-origin: 50% 30%; }
.card-hero.card-side .card-visual { width: 14rem; height: 14rem; font-size: 7rem; }
.card-full { --card-art-ink: #f2eee8; --card-art-ink-soft: rgba(242, 238, 232, 0.8); --card-art-shade: rgba(10, 8, 12, 0.9); isolation: isolate; min-height: 26rem; align-content: end; overflow: hidden; color: var(--card-art-ink); }
.card-full .card-visual { position: absolute; inset: 0; z-index: -1; width: auto; height: auto; border-radius: 0; font-size: 9rem; }
.card-full .card-visual::after { content: ""; position: absolute; inset: 0; background: linear-gradient(to bottom, transparent 35%, var(--card-art-shade) 82%); }
.card-full .card-fallback { align-items: start; padding-top: 2rem; background: radial-gradient(circle at 50% 30%, color-mix(in srgb, var(--card-accent, var(--pico-primary)) 45%, var(--pico-card-background-color)), var(--pico-card-background-color) 75%); }
.card-full :is(.card-subtitle, .card-meta) { color: var(--card-art-ink-soft); }
.card-art .card-fallback { background: color-mix(in srgb, var(--card-accent, var(--pico-primary)) 22%, var(--pico-card-background-color)); }
.card-face .card-visual img { transform: scale(1.6); transform-origin: 50% 28%; }
:where(.badge-row) { display: flex; flex-wrap: wrap; gap: 0.3rem; margin: 0.4rem 0 1rem; }
.pager { display: flex; align-items: center; gap: 0.75rem; }
:where(.pager-position) { color: var(--pico-muted-color); font-variant-numeric: tabular-nums; }
`
