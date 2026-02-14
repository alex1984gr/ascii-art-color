package tests

import (
	"bytes"
	"strings"
	"testing"

	"ascii-art/pipeline"
)

func TestPipeline_Run_FullFlow(t *testing.T) {
	// Provide arguments as they would be from command line
	args := []string{
		"--font=standard",
		"--color=red",
		"Hello",
	}

	// Use bytes.Buffer instead of stdout for testing
	var out bytes.Buffer

	// Run the pipeline with test arguments
	exitCode := pipeline.Run(args, &out)
	// Verify the pipeline succeeded (exit code 0)
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	// Get the output as a string
	output := out.String()

	// Check that output contains the ANSI code for red color
	if !strings.Contains(output, "\033[31m") {
		t.Errorf("output does not contain red ANSI code: %q", output)
	}

	// Check that output is not empty
	if len(output) == 0 {
		t.Errorf("output is empty")
	}
}

func TestPipeline_Run_InvalidFont(t *testing.T) {
	args := []string{
		"--font=nonexistent",
		"Test",
	}

	var out bytes.Buffer
	exitCode := pipeline.Run(args, &out)
	if exitCode == 0 {
		t.Fatalf("expected non-zero exit code for invalid font")
	}
}

func TestPipeline_Run_NoInput(t *testing.T) {
	args := []string{}

	var out bytes.Buffer
	exitCode := pipeline.Run(args, &out)
	if exitCode == 0 {
		t.Fatalf("expected non-zero exit code for empty input")
	}
}
