package constant

const InlineStyle = `
.summary-cards { display: flex; gap: 1rem; margin-bottom: 2rem; flex-wrap: wrap; }
.summary-cards article { flex: 1; min-width: 150px; text-align: center; }
table { margin-top: 1rem; }
.filter-form { margin-bottom: 1rem; }
.filter-form .grid { align-items: end; }
.success { background: #d4edda; color: #155724; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
.clickable-row { cursor: pointer; }
.clickable-row:hover { background: var(--pico-table-row-stripped-background-color); }
.detail-row td { padding: 0; }
.detail-content { padding: 1rem; }
.detail-content pre { white-space: pre-wrap; word-break: break-word; margin: 0.5rem 0; }
.detail-actions { display: flex; gap: 0.5rem; margin-top: 1rem; }
.detail-actions button { width: auto; margin-bottom: 0; }
.edit-form { padding: 1rem; }
.edit-form button { width: auto; }
`
