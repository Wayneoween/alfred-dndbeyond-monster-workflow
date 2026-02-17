package main

import (
	"os"
	"testing"
)

// TestMain sets up the test environment
func TestMain(m *testing.M) {
	// Set up minimal Alfred environment variables for testing
	// These are required by the awgo library
	if os.Getenv("alfred_workflow_bundleid") == "" {
		os.Setenv("alfred_workflow_bundleid", "net.deanishe.awgo.test")
	}
	if os.Getenv("alfred_workflow_cache") == "" {
		os.Setenv("alfred_workflow_cache", "/tmp/alfred-test-cache")
	}
	if os.Getenv("alfred_workflow_data") == "" {
		os.Setenv("alfred_workflow_data", "/tmp/alfred-test-data")
	}
	if os.Getenv("alfred_workflow_version") == "" {
		os.Setenv("alfred_workflow_version", "1.0.0")
	}
	if os.Getenv("alfred_workflow_name") == "" {
		os.Setenv("alfred_workflow_name", "alfred-dndbeyond-monster-workflow")
	}

	// Create test directories
	os.MkdirAll("/tmp/alfred-test-cache", 0755)
	os.MkdirAll("/tmp/alfred-test-data", 0755)

	// Run tests
	code := m.Run()

	// Clean up
	os.RemoveAll("/tmp/alfred-test-cache")
	os.RemoveAll("/tmp/alfred-test-data")

	os.Exit(code)
}
