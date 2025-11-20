const console = {
  log(...args) {
    globalThis.__dispatch("console", ...args);
  },
};

function close() {
  globalThis.__dispatch("close");
}

function fetch(url, options = {}) {
  const optionsJSON = JSON.stringify(options);
  const responseJSON = globalThis.__dispatch("fetch", url, optionsJSON);
  const response = JSON.parse(responseJSON);
  
  return {
    ok: response.ok,
    status: response.status,
    statusText: response.statusText,
    headers: response.headers,
    url: response.url,
    text: () => response.body,
    json: () => JSON.parse(response.body),
  };
}

const Kimera = {
  readFile(filePath) {
    return globalThis.__dispatch("readFile", filePath);
  },
  writeFile(filePath, content) {
    return globalThis.__dispatch("writeFile", filePath, content);
  },
};

globalThis = {
  ...globalThis,
  console,
  close,
  fetch,
  Kimera,
};

globalThis.window = globalThis;
