package main

import (
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Backup original args and restore after test
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Set os.Args to simulate command line arguments
	os.Args = []string{"main.go", "-str", `
{
  "a": "11",
"b": "22"
}
`, "-json"}

	main()
}
