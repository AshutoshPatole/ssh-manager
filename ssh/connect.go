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

	// Set your desired prompt colors using properly escaped ANSI escape codes
	prompt := "\\[\\033[1;31m\\]\\[\\033[42m\\]\\u@\\h \\[\\033[0;36m\\]\\w\\[\\033[0m\\]\\$"
	fmt.Println(prompt)

	// Construct a single SSH command to set the new PS1 configuration in ~/.bashrc
	exec.Command(
		"ssh",
		user+"@"+server,
		fmt.Sprintf(`sed -i '/^export PS1=/d' ~/.bashrc && echo 'export PS1="%s"' >> ~/.bashrc`, prompt),
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
