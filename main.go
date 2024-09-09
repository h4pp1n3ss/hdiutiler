package main

import (
	"fmt"
	"hdiutiler/utils"
	"os"
	"os/exec"

	"github.com/h4pp1n3ss/dmgutils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run hdiutiler.go <file-path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Step 1. Verify file integrity using hdiutil
	fmt.Println("[*] Starting integrity check ")

	// 1.1 Define the hdiutil command with 'verify' action
	cmd := exec.Command("hdiutil", "verify", filePath)

	fmt.Println("[!] Executing command: ", cmd.String())

	// 1.2 Run the command and capture both stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Print error message and output if command fails
		fmt.Printf("Error executing command: %s\n", err)
		os.Exit(1)
	}
	// 1.3 Check if word  "VALID" is in the output
	isvalid := utils.ContainsWord(string(output), "VALID")

	if isvalid {
		out := fmt.Sprintf("[!] hdiutil: verify checksum of %s is VALID", filePath)
		fmt.Println(out)
	}

	mountPoint := "/Volumes/RandomMountPointName"
	// Step 1: Mount the DMG file
	errno := dmgutils.MountDMG(filePath, mountPoint)

	if errno != nil {
		fmt.Printf("[-] Failed to mount DMG: %v\n", err)
		return
	}

	// Step 2: Locate the application
	appPath, err := dmgutils.FindApplication(mountPoint)
	if err != nil {
		fmt.Printf("[-] Failed to find application: %v\n", err)
		return
	}

	// Step 3: Verify code signing
	fmt.Println("[*] Starting mimics what Gatekeeper does to check your app")
	valid, err := dmgutils.VerifyCodeSigning(appPath)
	if err != nil {
		fmt.Printf("[-] Code signing verification failed: %v\n", err)
		return
	}

	if valid {
		fmt.Println("[!] Application is properly signed and trusted.")
	} else {
		fmt.Println("[-] Application is not properly signed or trusted.")
	}

	defer dmgutils.UnmountDMG(mountPoint) // Ensure unmounting after processing

}
