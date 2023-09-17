package ssh

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/TwiN/go-color"
)

func Connect(server, user string) {
	home, _ := os.UserHomeDir()

	privKey := home + "/.ssh/id_ed25519"
	// privKeyPath := fmt.Sprintf("-i %s", privKey)

	_, err := os.ReadFile(privKey)

	if err != nil {
		fmt.Println(color.InRed("Failed to read private key: " + err.Error()))
		return
	}

	sshCommand := exec.Command("ssh", user+"@"+server)

	// Set the standard input, output, and error streams to the current process's streams
	sshCommand.Stdin = os.Stdin
	sshCommand.Stdout = os.Stdout
	sshCommand.Stderr = os.Stderr

	ssh_err := sshCommand.Run()
	if ssh_err != nil {
		fmt.Println("Failed to run SSH command:", err)
	}
}
