package main

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Get the latest tag from git
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	out, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting latest tag: %s", err)
		os.Exit(1)
	}

	// Parse the tag into a semver.Version
	versionStr := strings.TrimSpace(string(out))
	if versionStr == "" {
		versionStr = "0.1.0" // Default to "v0.1.0" if there are no tags
	} else {
		versionStr = strings.TrimPrefix(versionStr, "v") // Remove the 'v' prefix if it exists
	}

	version, err := semver.NewVersion(versionStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing version: %s", err)
		os.Exit(1)
	}

	// Check the command line arguments for "break" or "feature"
	var newVersion semver.Version
	if len(os.Args) > 1 {
		if strings.Contains(os.Args[1], "break") {
			newVersion = version.IncMajor()
		} else if strings.Contains(os.Args[1], "feature") {
			newVersion = version.IncMinor()
		} else {
			newVersion = version.IncPatch()
		}
	} else {
		newVersion = version.IncPatch()
	}

	// Print the new version to stdout
	fmt.Print("v" + newVersion.String())
}
