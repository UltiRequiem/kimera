// HTTP server with route handling
console.log("Creating HTTP server with routes...");

const server = Kimera.createServer((request) => {
  console.log(`${request.method} ${request.path}`);
  
  // Route handling
  if (request.path === "/") {
    return {
      status: 200,
      headers: { "Content-Type": "text/html" },
      body: "<h1>Welcome to Kimera Server</h1><p>Try /api or /about</p>",
    };
  }
  
  if (request.path === "/api") {
    return {
      status: 200,
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ message: "API endpoint", version: "1.0" }),
    };
  }
  
  if (request.path === "/about") {
    return {
      status: 200,
      headers: { "Content-Type": "text/plain" },
      body: "Kimera.js HTTP Server - A minimal JavaScript runtime",
    };
  }
  
  // 404 Not Found
  return {
    status: 404,
    headers: { "Content-Type": "text/plain" },
    body: "Not Found",
  };
});

console.log("Starting server on port 8080...");
console.log("Try: http://localhost:8080/");
console.log("Try: http://localhost:8080/api");
console.log("Try: http://localhost:8080/about");
server.listen(8080);
