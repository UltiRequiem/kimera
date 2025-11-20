// Test that network operations are denied without --net flag
try {
  const response = fetch("https://example.com");
  console.log("ERROR: Should have denied network access");
  close();
} catch (error) {
  if (error.toString().includes("network access denied")) {
    console.log("PASS: Network access correctly denied");
  } else {
    console.log("ERROR: Wrong error message:", error);
    close();
  }
}
