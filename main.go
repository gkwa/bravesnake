package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Get the path of the executable file
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// Get the directory of the icon file
	iconDir := filepath.Dir(exePath)
	iconPath := filepath.Join(iconDir, "17902_power_shutdown_icon.png")

	// Convert the icon file to .ico format using rsrc
	cmd := exec.Command("rsrc", "-manifest", iconPath)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Rename the generated .syso file to icon.syso
	err = os.Rename("rsrc.syso", "icon.syso")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Icon converted successfully!")
}
