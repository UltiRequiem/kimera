// Test fetching text content
// Example usage:
// const response = fetch("https://example.com/page");
// const text = response.text();
// console.log("Content: " + text);

try {
  const response = fetch("http://localhost:8080/get");
  console.log("Status: " + response.status);
  console.log("OK: " + response.ok);
  
  const text = response.text();
  console.log("Response length: " + text.length);
  console.log("Text content: " + text);
  console.log("Text response successful!");
} catch (error) {
  console.log("Error: " + error);
}
