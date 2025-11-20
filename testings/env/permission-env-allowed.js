// Test that environment variable operations work with --env flag
try {
  // Test getEnv
  const path = Kimera.getEnv("PATH");
  if (path && path.length > 0) {
    console.log("PASS: getEnv works with --env flag");
  } else {
    console.log("ERROR: PATH should not be empty");
    close();
  }
  
  // Test setEnv
  Kimera.setEnv("KIMERA_TEST_VAR", "test_value");
  const testValue = Kimera.getEnv("KIMERA_TEST_VAR");
  if (testValue === "test_value") {
    console.log("PASS: setEnv works with --env flag");
  } else {
    console.log("ERROR: setEnv did not set the value correctly");
    close();
  }
} catch (error) {
  console.log("ERROR: Should have allowed environment access:", error);
  close();
}
