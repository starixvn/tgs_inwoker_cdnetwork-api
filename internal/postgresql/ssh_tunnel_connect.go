package postgresql

import (
	"cdnetwork/pkg/log"
	"cdnetwork/pkg/util"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

// PublicKeyFile reads and parses a private key file for SSH authentication
func PublicKeyFile(file string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(file)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Failed to read private key file: %v", err))
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Failed to parse private key: %v", err))
	}

	return ssh.PublicKeys(signer)
}

// connectSSH connects to an SSH server and returns an SSH client
func connectSSH(config util.SshTunnelConfig) (*ssh.Client, error) {
	// SSH configuration
	sshConfig := &ssh.ClientConfig{
		User: config.SshUser,
		Auth: []ssh.AuthMethod{
			PublicKeyFile(config.SshKeyFile),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: Insecure, use only for testing!
	}

	// Connect to SSH server
	sshConn, err := ssh.Dial("tcp", config.SshHost+":"+config.SshPort, sshConfig)
	if err != nil {
		return nil, err
	}
	return sshConn, nil
}
