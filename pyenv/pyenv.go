package pyenv

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var versionRegex = regexp.MustCompile(`\d+.\d+.\d+`)

func SetPythonChain(version string) {
	installCmd := exec.Command("pyenv", "install", version)
	activateCmd := exec.Command("pyenv", "local", version)
	installCmd.Stdout = os.Stdout
	activateCmd.Stdout = os.Stdout

	installCmd.Run()
	activateCmd.Run()
}

func GetAvailableVersions() []string {
	versions := exec.Command("pyenv", "versions")

	output, err := versions.CombinedOutput()
	if err != nil {
		fmt.Println("Error getting available versions:", err)
		return []string{}
	}

	outputStr := string(output)
	lines := strings.Split(strings.TrimSpace(outputStr), "\n")

	var cleanVersions []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		version := versionRegex.FindStringSubmatch(line)
		if len(version) > 0 {
			cleanVersions = append(cleanVersions, version[0])
		}
	}

	return cleanVersions
}

func ExecuteCommand(command []string) error {
	if len(command) > 0 {
		cmd := exec.Command("pyenv", "exec")
		cmd.Args = append(cmd.Args, command...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			return fmt.Errorf("error executing command: %v", err)
		}
		return nil
	}
	return fmt.Errorf("no command provided")
}
