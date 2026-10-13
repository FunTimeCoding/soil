const WAIT_TIMEOUT = 30000;
const POLL_INTERVAL = 100;
const NAVIGATION_GRACE = 1000;

function delay(milliseconds) {
  return new Promise(function (resolve) {
    setTimeout(resolve, milliseconds);
  });
}

function waitForNavigation(tabId, timeout) {
  return new Promise(function (resolve) {
    let done = false;

    function finish(value) {
      if (done) {
        return;
      }

      done = true;
      browser.tabs.onUpdated.removeListener(listen);
      resolve(value);
    }

    function listen(id, change) {
      if (id === tabId && change.status === "complete") {
        finish(true);
      }
    }

    browser.tabs.onUpdated.addListener(listen);
    setTimeout(function () {
      finish(false);
    }, timeout);
  });
}

async function waitForLoad(tabId, timeout) {
  const tab = await browser.tabs.get(tabId);

  if (tab.status === "complete") {
    return true;
  }

  return waitForNavigation(tabId, timeout);
}

function waitForLoadStart(tabId, timeout) {
  return new Promise(function (resolve) {
    let done = false;

    function finish(value) {
      if (done) {
        return;
      }

      done = true;
      browser.tabs.onUpdated.removeListener(listen);
      resolve(value);
    }

    function listen(id, change) {
      if (id === tabId && change.status === "loading") {
        finish(true);
      }
    }

    browser.tabs.onUpdated.addListener(listen);
    setTimeout(function () {
      finish(false);
    }, timeout);
  });
}

async function waitForSettle(tabId, timeout) {
  const tab = await browser.tabs.get(tabId);

  if (tab.status === "loading") {
    return waitForNavigation(tabId, timeout);
  }

  if (await waitForLoadStart(tabId, NAVIGATION_GRACE)) {
    return waitForNavigation(tabId, timeout);
  }

  return true;
}

async function waitForSelector(tabId, selector, timeout) {
  const deadline = Date.now() + timeout;
  const code = "!!document.querySelector(" + JSON.stringify(selector) + ")";

  for (;;) {
    try {
      const results = await browser.tabs.executeScript(tabId, { code: code });

      if (results && results[0]) {
        return true;
      }
    } catch (error) {
      // mid-navigation, executeScript is unavailable until the load settles
    }

    if (Date.now() >= deadline) {
      return false;
    }

    await delay(POLL_INTERVAL);
  }
}

function elementScript(selector, body) {
  return (
    "(function() {" +
    "var el = document.querySelector(" +
    JSON.stringify(selector) +
    ");" +
    "if (!el) { return null; }" +
    body +
    "return {selector: " +
    JSON.stringify(selector) +
    ", tag: el.tagName.toLowerCase()" +
    ', text: (el.value || el.innerText || "").trim().slice(0, 200)};' +
    "})()"
  );
}

async function runElementScript(tabId, selector, body) {
  if (!selector) {
    throw new Error("selector is required");
  }

  const results = await browser.tabs.executeScript(tabId, {
    code: elementScript(selector, body),
  });
  const element = results && results.length ? results[0] : null;

  if (!element) {
    throw new Error("selector matched nothing: " + selector);
  }

  return element;
}

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
  const created = await browser.tabs.create({
    url: params.url || "about:blank",
    active: true,
  });
  await waitForLoad(created.id, WAIT_TIMEOUT);
  const tab = await browser.tabs.get(created.id);

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
  const settled = waitForNavigation(tabId, WAIT_TIMEOUT);

  await browser.tabs.update(tabId, { url: params.url });
  await settled;
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

async function waitFor(params) {
  const tabId = params.tab_id;
  const timeout = params.timeout || WAIT_TIMEOUT;
  await waitForSettle(tabId, timeout);

  if (params.selector) {
    const found = await waitForSelector(tabId, params.selector, timeout);

    if (!found) {
      throw new Error("selector not found within timeout: " + params.selector);
    }
  }

  const tab = await browser.tabs.get(tabId);

  return {
    tab_id: tab.id,
    window_id: tab.windowId,
    url: tab.url || "",
    title: tab.title || "",
    active: tab.active,
    status: tab.status || "",
    index: tab.index,
    group_id: tab.groupId || -1,
  };
}

async function evaluate(params) {
  if (!params.expression) {
    throw new Error("expression is required");
  }

  const results = await browser.tabs.executeScript(params.tab_id, {
    code: "(function(){ return (" + params.expression + ") })()",
  });
  const value = results && results.length ? results[0] : null;

  return {
    tab_id: params.tab_id,
    value: value === undefined ? null : value,
  };
}

async function click(params) {
  return runElementScript(
    params.tab_id,
    params.selector,
    'el.scrollIntoView({block: "center"}); el.click();'
  );
}

async function fill(params) {
  const value = JSON.stringify(params.value || "");
  let body = "el.focus();";

  if (params.direct) {
    body +=
      "el.value = " +
      value +
      ';el.dispatchEvent(new Event("input", {bubbles: true}));' +
      'el.dispatchEvent(new Event("change", {bubbles: true}));';
  } else {
    body +=
      'document.execCommand("selectAll", false, null);' +
      'document.execCommand("insertText", false, ' +
      value +
      ");";
  }

  return runElementScript(params.tab_id, params.selector, body);
}

async function screenshot(params) {
  await browser.tabs.update(params.tab_id, { active: true });
  const tab = await browser.tabs.get(params.tab_id);
  const encoded = await browser.tabs.captureVisibleTab(tab.windowId, {
    format: "png",
  });

  return {
    tab_id: tab.id,
    url: tab.url || "",
    title: tab.title || "",
    content: encoded.slice(encoded.indexOf(",") + 1),
  };
}

async function downloadFile(params) {
  if (!params.url) {
    throw new Error("url is required");
  }

  const options = {
    url: params.url,
    saveAs: false,
    conflictAction: "uniquify",
  };

  if (params.filename) {
    options.filename = params.filename;
  }

  const identifier = await browser.downloads.download(options);

  return { download_id: identifier };
}

async function downloadStatus(params) {
  const query = {};

  if (params.download_id) {
    query.id = params.download_id;
  } else {
    query.limit = 10;
    query.orderBy = ["-startTime"];
  }

  const items = await browser.downloads.search(query);

  return items.map(function (item) {
    return {
      download_id: item.id,
      url: item.url || "",
      filename: item.filename || "",
      state: item.state || "",
      received: item.bytesReceived || 0,
      total: item.totalBytes || 0,
      error: item.error || "",
    };
  });
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
  wait_for: waitFor,
  evaluate: evaluate,
  click: click,
  fill: fill,
  screenshot: screenshot,
  download_file: downloadFile,
  download_status: downloadStatus,
};

async function dispatch(method, params) {
  const handler = methods[method];

  if (!handler) {
    throw new Error("unknown method: " + method);
  }

  return handler(params || {});
}
