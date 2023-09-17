package ssh

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/TwiN/go-color"
)

func Connect(server, user, environment string) {
	home, _ := os.UserHomeDir()

	privKey := home + "/.ssh/id_ed25519"
	// privKeyPath := fmt.Sprintf("-i %s", privKey)

	_, err := os.ReadFile(privKey)

	if err != nil {
		fmt.Println(color.InRed("Failed to read private key: " + err.Error()))
		return
	}
	var promptColor string
	// set bash prompt colors
	if environment == "prd" {
		promptColor = "\\[\\033[1;31m\\]\\[\\033m\\]\\u@\\h \\[\\033[1;36m\\]\\w\\[\\033[0m\\]\\$ "
	} else if environment == "dev" {
		promptColor = "\\[\\033[1;32m\\]\\[\\033m\\]\\u@\\h \\[\\033[1;36m\\]\\w\\[\\033[0m\\]\\$ "
	} else if environment == "uat" {
		promptColor = "\\[\\033[1;34m\\]\\[\\033m\\]\\u@\\h \\[\\033[1;36m\\]\\w\\[\\033[0m\\]\\$ "
	} else if environment == "sit" {
		promptColor = "\\[\\033[1;34m\\]\\[\\033m\\]\\u@\\h \\[\\033[1;36m\\]\\w\\[\\033[0m\\]\\$ "
	} else if environment == "ppd" {
		promptColor = "\\[\\033[1;33m\\]\\[\\033m\\]\\u@\\h \\[\\033[1;36m\\]\\w\\[\\033[0m\\]\\$ "
	}
	// Construct a single SSH command to set the new PS1 configuration in ~/.bashrc
	exec.Command(
		"ssh",
		user+"@"+server,
		fmt.Sprintf(`sed -i '/^export PS1=/d' ~/.bashrc && echo 'export PS1="%s"' >> ~/.bashrc`, promptColor),
	).Run()

	// Modify the SSH command to set the prompt colors
	sshCommand := exec.Command("ssh", user+"@"+server)

	// Set the standard input, output, and error streams to the current process's streams
	sshCommand.Stdin = os.Stdin
	sshCommand.Stdout = os.Stdout
	sshCommand.Stderr = os.Stderr

	ssh_err := sshCommand.Run()
	if ssh_err != nil {
		fmt.Println("Failed to run SSH command:", ssh_err)
	}
}
