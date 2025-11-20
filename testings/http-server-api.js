// REST API server example combining HTTP server with other features
console.log("Creating REST API server...");

// In-memory data store
const users = [
  { id: 1, name: "Alice", email: "alice@example.com" },
  { id: 2, name: "Bob", email: "bob@example.com" },
];

let nextId = 3;

const server = Kimera.createServer((request) => {
  console.log(`${request.method} ${request.path}`);
  
  const headers = {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  };
  
  // GET /users - List all users
  if (request.path === "/users" && request.method === "GET") {
    return {
      status: 200,
      headers,
      body: JSON.stringify({ users }),
    };
  }
  
  // POST /users - Create a new user
  if (request.path === "/users" && request.method === "POST") {
    try {
      const newUser = JSON.parse(request.body);
      const user = {
        id: nextId++,
        name: newUser.name,
        email: newUser.email,
      };
      users.push(user);
      
      return {
        status: 201,
        headers,
        body: JSON.stringify({ user, message: "User created" }),
      };
    } catch (error) {
      return {
        status: 400,
        headers,
        body: JSON.stringify({ error: "Invalid JSON" }),
      };
    }
  }
  
  // GET /users/:id - Get user by ID
  if (request.path.startsWith("/users/") && request.method === "GET") {
    const id = parseInt(request.path.split("/")[2]);
    const user = users.find(u => u.id === id);
    
    if (user) {
      return {
        status: 200,
        headers,
        body: JSON.stringify({ user }),
      };
    } else {
      return {
        status: 404,
        headers,
        body: JSON.stringify({ error: "User not found" }),
      };
    }
  }
  
  // GET /health - Health check
  if (request.path === "/health") {
    return {
      status: 200,
      headers,
      body: JSON.stringify({ 
        status: "healthy",
        users: users.length,
        timestamp: Date.now(),
      }),
    };
  }
  
  // Root endpoint
  if (request.path === "/") {
    return {
      status: 200,
      headers: { "Content-Type": "text/html" },
      body: `
        <h1>Kimera REST API</h1>
        <ul>
          <li>GET /users - List all users</li>
          <li>POST /users - Create a user</li>
          <li>GET /users/:id - Get user by ID</li>
          <li>GET /health - Health check</li>
        </ul>
      `,
    };
  }
  
  // 404 for unknown routes
  return {
    status: 404,
    headers,
    body: JSON.stringify({ error: "Not Found" }),
  };
});

console.log("API server running on port 8080");
console.log("Try: curl http://localhost:8080/");
console.log("Try: curl http://localhost:8080/users");
console.log("Try: curl http://localhost:8080/health");
console.log("Try: curl -X POST -H 'Content-Type: application/json' -d '{\"name\":\"Charlie\",\"email\":\"charlie@example.com\"}' http://localhost:8080/users");

server.listen(8080);
