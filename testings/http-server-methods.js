// HTTP server demonstrating different HTTP methods
console.log("Creating HTTP server with method handling...");

const server = Kimera.createServer((request) => {
  console.log(`${request.method} ${request.path}`);
  
  if (request.path === "/data") {
    if (request.method === "GET") {
      return {
        status: 200,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ data: "Sample data", items: [1, 2, 3] }),
      };
    }
    
    if (request.method === "POST") {
      return {
        status: 201,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ 
          message: "Data created", 
          received: request.body 
        }),
      };
    }
    
    if (request.method === "PUT") {
      return {
        status: 200,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ 
          message: "Data updated",
          received: request.body 
        }),
      };
    }
    
    if (request.method === "DELETE") {
      return {
        status: 200,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ message: "Data deleted" }),
      };
    }
    
    // Method not allowed
    return {
      status: 405,
      headers: { "Content-Type": "text/plain" },
      body: "Method Not Allowed",
    };
  }
  
  return {
    status: 200,
    headers: { "Content-Type": "text/plain" },
    body: "Try POST/GET/PUT/DELETE to /data",
  };
});

console.log("Starting server on port 8080...");
console.log("Test with: curl http://localhost:8080/data");
console.log("Test with: curl -X POST -d 'test' http://localhost:8080/data");
console.log("Test with: curl -X PUT -d 'update' http://localhost:8080/data");
console.log("Test with: curl -X DELETE http://localhost:8080/data");
server.listen(8080);
