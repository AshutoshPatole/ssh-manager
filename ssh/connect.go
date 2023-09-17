package ssh

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/TwiN/go-color"
	"golang.org/x/crypto/ssh"
)

func Connect(server, user string) {

	home, _ := os.UserHomeDir()

	privKey := home + "/.ssh/id_ed25519"
	privKeyBytes, err := os.ReadFile(privKey)

	if err != nil {
		fmt.Println(color.InRed("Failed to read private key: " + err.Error()))
		return
	}

	signer, err := ssh.ParsePrivateKey(privKeyBytes)

	if err != nil {
		fmt.Println(color.InRed("Failed to parse private key: " + err.Error()))
		return
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},

		// TODO: fix insecureignore host callback method
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// optional tcp connect timeout
		Timeout: 5 * time.Second,
	}
	client, err := ssh.Dial("tcp", server+":"+PORT, config)
	if err != nil {
		fmt.Println(color.InRed(err.Error()))
		return
	}
	defer client.Close()

	// start session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println(color.InRed(err.Error()))
		return
	}
	defer session.Close()

	// // setup standard out and error
	// // uses writer interface
	// session.Stdout = os.Stdout
	// session.Stderr = os.Stderr
	var stdin io.WriteCloser
	var stdout, stderr io.Reader

	stdin, err = session.StdinPipe()
	if err != nil {
		fmt.Println(err.Error())
	}

	stdout, err = session.StdoutPipe()
	if err != nil {
		fmt.Println(err.Error())
	}

	stderr, err = session.StderrPipe()
	if err != nil {
		fmt.Println(err.Error())
	}
	// if err := session.Run("hostname -f "); err != nil {
	// 	fmt.Print("Could not find host")
	// }
	// defer session.Close()

	// err = session.Shell()
	// if err != nil {
	// 	fmt.Println(color.InRed(err.Error()))
	// 	return
	// }

	wr := make(chan []byte, 100)
	go func() {
		scanner := bufio.NewScanner(stdout)
		for {
			if tkn := scanner.Scan(); tkn {
				rcv := scanner.Bytes()

				raw := make([]byte, len(rcv))
				copy(raw, rcv)

				fmt.Println(string(raw))
			} else {
				if scanner.Err() != nil {
					fmt.Println(scanner.Err())
				} else {
					fmt.Println("io.EOF")
				}
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case d := <-wr:
				_, err := stdin.Write(d)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for {
			if tkn := scanner.Scan(); tkn {
				rcv := scanner.Bytes()

				raw := make([]byte, len(rcv))
				copy(raw, rcv)

				fmt.Println(string(raw))
			} else {
				if scanner.Err() != nil {
					fmt.Println(scanner.Err())
				} else {
					fmt.Println("io.EOF")
				}
				return
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	session.Shell()

	for {
		fmt.Print("$")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		wr <- []byte(text + "\n")
	}

}
