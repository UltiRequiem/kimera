// Comprehensive fetch API example for Kimera.js
// This example demonstrates the complete fetch API functionality

console.log("=== Kimera.js Fetch API Examples ===\n");

// Example 1: Simple GET request
console.log("1. Simple GET Request:");
console.log("   const response = fetch('https://api.example.com/data');");
console.log("   console.log('Status: ' + response.status);");
console.log("   const data = response.json();\n");

// Example 2: POST request with JSON body
console.log("2. POST Request with JSON:");
console.log("   const response = fetch('https://api.example.com/create', {");
console.log("     method: 'POST',");
console.log("     body: JSON.stringify({ name: 'John', age: 30 }),");
console.log("     headers: { 'Content-Type': 'application/json' }");
console.log("   });");
console.log("   console.log('Created: ' + response.status);\n");

// Example 3: Using text() method
console.log("3. Getting Text Response:");
console.log("   const response = fetch('https://example.com/page');");
console.log("   const html = response.text();");
console.log("   console.log('Page content: ' + html);\n");

// Example 4: Checking response status
console.log("4. Checking Response Status:");
console.log("   const response = fetch('https://api.example.com/data');");
console.log("   if (response.ok) {");
console.log("     console.log('Success!');");
console.log("   } else {");
console.log("     console.log('Error: ' + response.status);");
console.log("   }\n");

// Example 5: Custom headers
console.log("5. Custom Headers:");
console.log("   const response = fetch('https://api.example.com/data', {");
console.log("     headers: {");
console.log("       'Authorization': 'Bearer token123',");
console.log("       'Custom-Header': 'value'");
console.log("     }");
console.log("   });\n");

console.log("=== For more examples, see testings/fetch-*.js ===");
