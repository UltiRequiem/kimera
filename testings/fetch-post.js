// Test POST request with body
// Example usage:
// const response = fetch("https://api.example.com/data", {
//   method: "POST",
//   body: JSON.stringify({ key: "value" }),
//   headers: { "Content-Type": "application/json" }
// });
// console.log("Status: " + response.status);

try {
  const response = fetch("http://localhost:8080/post", {
    method: "POST",
    body: JSON.stringify({
      title: "Test Post from Kimera",
      body: "This is a test post",
      userId: 1
    }),
    headers: {
      "Content-Type": "application/json"
    }
  });
  
  console.log("Status: " + response.status);
  console.log("OK: " + response.ok);
  
  const data = response.json();
  console.log("Server message: " + data.message);
  console.log("POST request successful!");
} catch (error) {
  console.log("Error: " + error);
}
