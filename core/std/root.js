const console = {
  log(...args) {
    globalThis.__dispatch("console", ...args);
  },
};

function close() {
  globalThis.__dispatch("close");
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
  Kimera,
};

globalThis.window = globalThis;
