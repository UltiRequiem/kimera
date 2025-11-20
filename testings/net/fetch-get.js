// Test basic GET request
// Example usage:
// const response = fetch("https://api.example.com/data");
// console.log("Status: " + response.status);
// console.log("OK: " + response.ok);
// const data = response.json();
// console.log("Data: " + JSON.stringify(data));

try {
  const response = fetch("https://joke.deno.dev");
  console.log("Status: " + response.status);
  console.log("OK: " + response.ok);

  const data = response.json();
  console.log("Message: " + data.message);
  console.log("Method: " + data.method);
  console.log(response.data);
  console.log("GET request successful!");
} catch (error) {
  console.log("Error: " + error);
}
