// Basic HTTP server example
console.log("Creating HTTP server...");

const server = Kimera.createServer((request) => {
  console.log(`${request.method} ${request.path}`);
  
  return {
    status: 200,
    headers: {
      "Content-Type": "text/plain",
    },
    body: "Hello from Kimera HTTP Server!",
  };
});

console.log("Starting server on port 8080...");
server.listen(8080);
