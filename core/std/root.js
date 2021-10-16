globalThis.console = {
  log: (...args) => {
    globalThis.__dispatch("console", ...args);
  },
};
