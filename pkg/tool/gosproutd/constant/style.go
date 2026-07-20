package constant

const InlineStyle = `
.seed-table { width: 100%; }
.seed-table th, .seed-table td { padding: 0.5rem 0.75rem; }
.position-cell { width: 3rem; text-align: center; }
.seed-table tbody.sortable tr { cursor: grab; }
.seed-table tbody.sortable tr:active { cursor: grabbing; }
.sortable-ghost { opacity: 0.4; background: var(--pico-primary-focus); }
.sortable-chosen { background: var(--pico-card-background-color); }
.drag-handle { cursor: grab; user-select: none; color: var(--pico-muted-color); }
`
