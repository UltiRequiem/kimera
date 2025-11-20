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

// Initialize server handlers registry
if (!globalThis.__serverHandlers) {
  globalThis.__serverHandlers = {};
}

let __handlerCounter = 0;

const Kimera = {
  readFile(filePath) {
    return globalThis.__dispatch("readFile", filePath);
  },
  writeFile(filePath, content) {
    return globalThis.__dispatch("writeFile", filePath, content);
  },
  getEnv(varName) {
    return globalThis.__dispatch("getEnv", varName);
  },
  setEnv(varName, value) {
    return globalThis.__dispatch("setEnv", varName, value);
  },
  createServer(handler) {
    // Generate unique handler ID
    const handlerID = `handler_${__handlerCounter++}`;
    
    // Store the handler function
    globalThis.__serverHandlers[handlerID] = handler;
    
    // Create server ID through dispatch
    const serverID = globalThis.__dispatch("httpCreateServer", handlerID);
    
    return {
      listen(port) {
        return globalThis.__dispatch("httpServerListen", serverID, String(port), handlerID);
      },
      close() {
        // Clean up handler when server closes
        delete globalThis.__serverHandlers[handlerID];
        return globalThis.__dispatch("httpServerClose", serverID);
      },
    };
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
