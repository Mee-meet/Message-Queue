package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func runGoFileInTerminal(filePath string) {
	var cmd *exec.Cmd

	// Choose the appropriate command based on the OS.
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", "go", "run", filePath)
	} else {
		log.Fatal("This platform is not supported. Only Windows is supported.")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// List of paths to the three Go files
	goFiles := []string{
		"D:/TEMP/notif/AES/Email_consumer/emailconsumer.go",
		"D:/TEMP/notif/AES/sms_consumer/smsconsumerr.go",
		"D:/TEMP/notif/AES/Whatsapp_consumer/whatsappconsumer.go",
	}

	// Run each Go file in a separate terminal
	for _, goFile := range goFiles {
		runGoFileInTerminal(goFile)
	}

	log.Println("All processes started.")
}
