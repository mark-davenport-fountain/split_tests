package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadJUnitXML(t *testing.T) {
	t.Run("successfully decodes valid junit XML", func(t *testing.T) {
		file, err := os.Open(filepath.Join("testdata", "valid_junit.xml"))
		if err != nil {
			t.Fatalf("Failed to open test fixture: %v", err)
		}
		defer file.Close()

		result := loadJUnitXML(file)

		if result == nil {
			t.Fatal("Expected non-nil result")
		}

		// Verify the parsed data
		if len(result.Testsuites) != 1 {
			t.Errorf("Expected 1 test suite, got %d", len(result.Testsuites))
		}

		suite := result.Testsuites[0]
		if len(suite.TestCases) != 2 {
			t.Errorf("Expected 2 test cases, got %d", len(suite.TestCases))
		}

		// Verify the test case details
		if suite.TestCases[0].Name != "successful_test" {
			t.Errorf("Expected test case name 'successful_test', got '%s'", suite.TestCases[0].Name)
		}

		if suite.TestCases[1].Name != "failed_test" {
			t.Errorf("Expected test case name 'failed_test', got '%s'", suite.TestCases[1].Name)
		}

		if suite.TestCases[1].File != "sample_spec_2.rb" {
			t.Errorf("Expected test case file 'sample_spec_2.rb', got '%s'", suite.TestCases[1].File)
		}

		if suite.TestCases[0].File != "sample_spec_1.rb" {
			t.Errorf("Expected test case file 'sample_spec_1.rb', got '%s'", suite.TestCases[0].File)
		}

	})
}
