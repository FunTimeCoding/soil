async function listTabs() {
  const tabs = await browser.tabs.query({});
  return tabs.map(tab => ({
    tab_id: tab.id,
    window_id: tab.windowId,
    url: tab.url || "",
    title: tab.title || "",
    active: tab.active,
    status: tab.status || "",
    index: tab.index,
    group_id: tab.groupId || -1,
  }));
}

async function readTab(params) {
  const tabId = params.tab_id;
  const raw = params.raw || false;
  const tab = await browser.tabs.get(tabId);
  let text = "";
  let mode = "raw";

  if (!raw) {
    try {
      await browser.tabs.executeScript(tabId, { file: "Readability.js" });
      const results = await browser.tabs.executeScript(tabId, {
        code: `
          (function() {
            var clone = document.cloneNode(true);
            var article = new Readability(clone).parse();
            if (article && article.textContent) {
              return article.textContent;
            }
            return null;
          })()
        `,
      });

      if (results[0]) {
        text = results[0];
        mode = "readability";
      }
    } catch (error) {
      // readability failed, fall through to innerText
    }
  }

  if (!text) {
    const results = await browser.tabs.executeScript(tabId, {
      code: "document.body.innerText",
    });
    text = results[0] || "";
    mode = "raw";
  }

  return {
    tab_id: tab.id,
    url: tab.url || "",
    title: tab.title || "",
    text: text,
    mode: mode,
  };
}

async function createTab(params) {
  const tab = await browser.tabs.create({
    url: params.url || "about:blank",
    active: true,
  });

  return {
    tab_id: tab.id,
    url: tab.url || "",
    title: tab.title || "",
    active: tab.active,
  };
}

async function closeTab(params) {
  const tabId = params.tab_id;
  const tab = await browser.tabs.get(tabId);
  const summary = {
    tab_id: tab.id,
    url: tab.url || "",
    title: tab.title || "",
  };

  await browser.tabs.remove(tabId);

  return summary;
}

async function navigate(params) {
  const tabId = params.tab_id;

  await browser.tabs.update(tabId, { url: params.url });
  const tab = await browser.tabs.get(tabId);

  return {
    tab_id: tab.id,
    url: tab.url || "",
    title: tab.title || "",
  };
}

async function navigateBack(params) {
  const tabId = params.tab_id;

  await browser.tabs.executeScript(tabId, {
    code: "history.back()",
  });

  const tab = await browser.tabs.get(tabId);

  return {
    tab_id: tab.id,
    url: tab.url || "",
    title: tab.title || "",
  };
}

async function groupTabs(params) {
  const tabIds = params.tab_ids;
  const options = { tabIds: tabIds };

  if (params.group_id && params.group_id > 0) {
    options.groupId = params.group_id;
  }

  const groupId = await browser.tabs.group(options);

  if (params.title || params.color) {
    const update = {};
    if (params.title) update.title = params.title;
    if (params.color) update.color = params.color;
    await browser.tabGroups.update(groupId, update);
  }

  return { group_id: groupId, tab_ids: tabIds };
}

async function ungroupTab(params) {
  const tabId = params.tab_id;
  await browser.tabs.ungroup(tabId);

  return { tab_id: tabId };
}

async function updateGroup(params) {
  const groupId = params.group_id;
  const update = {};

  if (params.title !== undefined) update.title = params.title;
  if (params.color !== undefined) update.color = params.color;
  if (params.collapsed !== undefined) update.collapsed = params.collapsed;

  await browser.tabGroups.update(groupId, update);

  return { group_id: groupId, ...update };
}

const methods = {
  list_tabs: listTabs,
  read_tab: readTab,
  create_tab: createTab,
  close_tab: closeTab,
  navigate: navigate,
  navigate_back: navigateBack,
  group_tabs: groupTabs,
  ungroup_tab: ungroupTab,
  update_group: updateGroup,
};

async function dispatch(method, params) {
  const handler = methods[method];

  if (!handler) {
    throw new Error("unknown method: " + method);
  }

  return handler(params || {});
}
