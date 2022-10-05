package myssh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Session *ssh.Session
	StdOut  bytes.Buffer
	StdErr  bytes.Buffer
}

func GetSSHClient(ip string, port int, username string, password string) (s *SSHClient, err error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", ip, port), config)
	if err != nil {
		return
	}
	session, err := client.NewSession()
	if err != nil {
		return
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	s = &SSHClient{
		Session: session,
		StdOut:  stdout,
		StdErr:  stderr,
	}
	session.Stdout = &s.StdOut
	session.Stderr = &s.StdErr
	return
}

func (s *SSHClient) sshRun(cmd string) (result string, err error) {
	err = s.Session.Run(cmd)
	if err != nil {
		return s.StdErr.String(), err
	}
	return s.StdOut.String(), nil
}
