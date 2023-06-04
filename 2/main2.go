package main2

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Paths to the input ICO file and the output executable
	icoPath := "17902_power_shutdown_icon.ico"
	exePath := "path/to/output.exe"

	// Generate the .syso file using rsrc
	err := generateResourceFile(icoPath, exePath)
	if err != nil {
		fmt.Println("Error generating resource file:", err)
		return
	}

	// Build the executable
	err = buildExecutable(exePath)
	if err != nil {
		fmt.Println("Error building executable:", err)
		return
	}

	fmt.Println("Executable with icon created successfully.")
}

func generateResourceFile(icoPath, exePath string) error {
	args := []string{
		"-ico", icoPath,
		"-o", exePath + ".syso",
	}

	cmd := exec.Command("rsrc", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to generate resource file: %v, output: %s", err, string(output))
	}

	return nil
}

func buildExecutable(exePath string) error {
	// Split the executable path to get the directory and the executable name
	dir, name := splitPath(exePath)

	args := []string{
		"build",
		"-o", exePath,
		"-ldflags", "-H=windowsgui",
	}

	// Change to the directory where the executable will be built
	err := os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	// Build the executable using the go command
	cmd := exec.Command("go", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to build executable: %v, output: %s", err, string(output))
	}

	return nil
}

func splitPath(path string) (string, string) {
	lastSlash := strings.LastIndex(path, "/")
	lastBackslash := strings.LastIndex(path, "\\")
	lastIndex := lastSlash
	if lastBackslash > lastSlash {
		lastIndex = lastBackslash
	}

	if lastIndex == -1 {
		return "", path
	}

	return path[:lastIndex], path[lastIndex+1:]
}
