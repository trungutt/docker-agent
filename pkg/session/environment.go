package session

import (
	"fmt"
	"runtime"
)

// getEnvironmentInfo returns formatted environment information including
// working directory, git repository status, and platform information
func getEnvironmentInfo(workingDir string) string {
	return fmt.Sprintf(`Here is useful information about the environment you are running in:
	<env>
	Working directory: %s
	Is directory a git repo: %s
	Operating System: %s
	CPU Architecture: %s
	</env>`, workingDir, boolToYesNo(isGitRepo(workingDir)), getOperatingSystem(), getArchitecture())
}

// boolToYesNo converts a boolean to "Yes" or "No" string
func boolToYesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

func getOperatingSystem() string {
	return operatingSystemName(runtime.GOOS)
}

// operatingSystemName maps a runtime.GOOS value to a human-readable OS name.
// Kept as a pure function so it can be tested independently of the running platform.
func operatingSystemName(goos string) string {
	switch goos {
	case "darwin":
		return "MacOS"
	case "windows":
		return "Windows"
	case "linux":
		return "Linux"
	default:
		return goos
	}
}

func getArchitecture() string {
	return architectureName(runtime.GOARCH)
}

// architectureName maps a runtime.GOARCH value to a human-readable architecture name.
// Kept as a pure function so it can be tested independently of the running platform.
func architectureName(goarch string) string {
	switch goarch {
	case "amd64":
		return "x64"
	case "arm64":
		return "arm64"
	default:
		return goarch
	}
}
