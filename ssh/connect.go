package ssh

import (
	"log"
	"os"
	"time"

	cConfig "github.com/AshutoshPatole/ssh-manager/config"
	"github.com/TwiN/go-color"
	"golang.org/x/crypto/ssh"
)

const PORT = "22"

func ConnectServer(server, user, password, group, env string) {

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
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
		log.Fatalln(color.InRed(err.Error()))
	}
	defer client.Close()

	// start session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalln(color.InRed(err))
	}
	defer session.Close()

	// setup standard out and error
	// uses writer interface
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	success := AddPubKeysToServer(session)

	cConfig.SaveServer(server, user, group, env, success)

}
