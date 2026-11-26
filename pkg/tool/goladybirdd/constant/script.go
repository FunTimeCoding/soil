package constant

const SnapshotScript = `() => {
  const implicit = {
    A: (e) => (e.hasAttribute("href") ? "link" : null),
    BUTTON: () => "button",
    SUMMARY: () => "button",
    INPUT: (e) => ({ checkbox: "checkbox", radio: "radio", submit: "button", button: "button", reset: "button", image: "button", search: "searchbox", range: "slider", hidden: null }[e.type] ?? "textbox"),
    TEXTAREA: () => "textbox",
    SELECT: () => "combobox",
    OPTION: () => "option",
    H1: () => "heading", H2: () => "heading", H3: () => "heading",
    H4: () => "heading", H5: () => "heading", H6: () => "heading",
    NAV: () => "navigation", MAIN: () => "main", HEADER: () => "banner",
    FOOTER: () => "contentinfo", ASIDE: () => "complementary",
    FORM: () => "form", DIALOG: () => "dialog",
    IMG: (e) => (e.getAttribute("alt") ? "img" : null),
    UL: () => "list", OL: () => "list", LI: () => "listitem",
    TABLE: () => "table", TR: () => "row", TH: () => "columnheader", TD: () => "cell",
    LABEL: () => "label",
  };
  const content = new Set(["link", "button", "heading", "label", "option", "cell", "columnheader", "checkbox", "radio", "tab", "menuitem"]);
  const unvalued = new Set(["checkbox", "radio", "submit", "button", "reset", "image", "file", "hidden", "password"]);
  const clip = (s, n) => {
    s = (s ?? "").replace(/\s+/g, " ").trim();
    return s.length > n ? s.slice(0, n - 1) + "…" : s;
  };
  const hidden = (e) => {
    if (e.hidden || e.getAttribute("aria-hidden") === "true") return true;
    const style = getComputedStyle(e);
    return style.display === "none" || style.visibility === "hidden";
  };
  const name = (e, role) => {
    const label = e.getAttribute("aria-label");
    if (label) return label;
    const by = e.getAttribute("aria-labelledby");
    if (by) return by.split(/\s+/).map((i) => document.getElementById(i)?.innerText ?? "").join(" ");
    if (e.labels && e.labels.length) return [...e.labels].map((l) => l.innerText).join(" ");
    const alt = e.getAttribute("alt");
    if (alt) return alt;
    if (content.has(role) && e.childElementCount < 8) {
      const text = e.innerText ?? e.textContent;
      if (text && text.trim()) return text;
    }
    return e.getAttribute("title") ?? e.getAttribute("placeholder") ?? "";
  };
  const nodes = [];
  const walk = (e, into, named) => {
    if (hidden(e)) return;
    const role = e.getAttribute("role") ?? implicit[e.tagName]?.(e) ?? null;
    let target = into;
    let quiet = named;
    if (role) {
      const item = { i: nodes.length, r: role, n: clip(name(e, role), 100), c: [] };
      const valued = e.tagName === "TEXTAREA" || e.tagName === "SELECT" || (e.tagName === "INPUT" && !unvalued.has(e.type));
      if (valued) item.v = clip(e.value, 100);
      nodes.push(e);
      into.push(item);
      target = item.c;
      quiet = named || (content.has(role) && item.n !== "");
    }
    for (const child of e.childNodes) {
      if (child.nodeType === 1) walk(child, target, quiet);
      else if (child.nodeType === 3 && !quiet && child.textContent.trim()) target.push({ r: "text", n: clip(child.textContent, 120), c: [] });
    }
  };
  const root = [];
  walk(document.body, root, false);
  return [JSON.stringify({ title: document.title, tree: root }), ...nodes];
}`

const (
	ClickScript  = "(e) => e.click()"
	FocusScript  = "(e) => { e.focus(); if (e.select) e.select(); }"
	DirectScript = "(e, v) => { e.value = v; e.dispatchEvent(new Event('input', { bubbles: true })); e.dispatchEvent(new Event('change', { bubbles: true })); }"
)
