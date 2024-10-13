package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Main test function to verify the pipeline
func TestDaggerPipeline(t *testing.T) {
	if shouldRunPipeline() {
		t.Log("Generated files not found. Running 'go run Daggerpipeline.go' to build the necessary artifacts.")
		runDaggerPipeline(t)
	} else {
		t.Log("Required artifacts already exist. Skipping 'go run Daggerpipeline.go'.")
	}

	// Always run verification tests regardless of whether we executed the pipeline or not
	verifyExecutableBuilds(t)
	verifyDockerTarExported(t)
	verifyTagNameAndVersion(t)
	verifyRunContainerScript(t)
}

// Function to determine whether to run Daggerpipeline.go
func shouldRunPipeline() bool {
	// Check if the Docker tar file or the executables exist
	_, tarErr := os.Stat("dagger-excuse-deno.tar")
	_, linuxErr := os.Stat("excuse-linux")
	_, macErr := os.Stat("excuse-mac")
	_, winErr := os.Stat("excuse-win.exe")

	// If any of the files do not exist, we need to run the pipeline
	return os.IsNotExist(tarErr) || os.IsNotExist(linuxErr) || os.IsNotExist(macErr) || os.IsNotExist(winErr)
}

// Function to run Daggerpipeline.go
func runDaggerPipeline(t *testing.T) {
	cmd := exec.Command("go", "run", "Daggerpipeline.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run Daggerpipeline.go: %v\nOutput: %s", err, string(output))
	}
}

func verifyExecutableBuilds(t *testing.T) {
	executables := []string{"excuse-linux", "excuse-mac", "excuse-win.exe"}
	for _, executable := range executables {
		if _, err := exec.Command("ls", executable).Output(); err != nil {
			t.Errorf("Expected executable %s not found", executable)
		}
	}
}

func verifyDockerTarExported(t *testing.T) {
	tarFileName := "dagger-excuse-deno.tar"
	if _, err := os.Stat(tarFileName); os.IsNotExist(err) {
		t.Errorf("Expected Docker tar file %s not found", tarFileName)
	}
}

func verifyTagNameAndVersion(t *testing.T) {
	tagName := "dagger-excuse-deno:latest"
	cmd := exec.Command("docker", "inspect", tagName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to inspect docker image: %v", err)
	}
	if !strings.Contains(string(output), tagName) {
		t.Errorf("Expected tag name %s not found", tagName)
	}
}

func verifyRunContainerScript(t *testing.T) {
	cmd := exec.Command("./run-container.sh")
	err := cmd.Run()
	if err != nil {
		t.Errorf("run-container.sh failed to execute: %v", err)
	}
}
