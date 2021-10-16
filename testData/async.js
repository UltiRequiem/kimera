const asyncSayHello = async (text) => {
  console.log(text);
};

(async () => await asyncSayHello("Hello World!"))();
