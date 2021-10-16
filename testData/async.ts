const asyncSayHello = async (text: string) => {
  console.log(text);
};

(async () => await asyncSayHello("Hello World!"))();
