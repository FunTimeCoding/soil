package constant

const InlineStyle = `
.seed-table { width: 100%; }
.seed-table th, .seed-table td { padding: 0.5rem 0.75rem; }
#seed-order tr { cursor: grab; }
#seed-order tr:active { cursor: grabbing; }
#seed-order td.drag-handle { width: 2rem; }
#seed-order td.numeric { width: 3rem; }
.sortable-ghost { opacity: 0.4; background: var(--pico-primary-focus); }
.sortable-chosen { background: var(--pico-card-background-color); }
.drag-handle { cursor: grab; user-select: none; color: var(--pico-muted-color); }
.seed-path, .counter-note { color: var(--pico-muted-color); }
.card-task .card-head h4 { font-size: 1.05rem; }
.card-task.card-hero .card-head h1 { font-size: 1.35rem; }
.card-footer .cluster:empty { display: none; }
`
