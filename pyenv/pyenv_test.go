package pyenv

import (
	"os/exec"
	"strings"
	"testing"
)

func TestGetAvailableVersions(t *testing.T) {
	versions := GetAvailableVersions()
	if len(versions) == 0 {
		t.Errorf("No versions found")
	}
}

func TestSetPythonChain(t *testing.T) {
	version := "3.5"
	SetPythonChain(version)

	cmd := exec.Command("pyenv", "exec", "python", "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}

	if !strings.Contains(string(output), version) {
		t.Errorf("Python version %s not found", version)
	}
}

func TestExecuteCommand(t *testing.T) {
	command := []string{"python", "--version"}
	err := ExecuteCommand(command)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
}
