package remote

import (
	"fmt"
	"log"
	"net"
	"publish/config"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var sshClient *ssh.Client

func GetSftpClient() (*sftp.Client, error) {
	var sftpClient *sftp.Client
	var err error
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}

func GetSession() (*ssh.Session, error) {
	var session *ssh.Session
	var err error
	if session, err = sshClient.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

func init() {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(config.Server.Password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User:            config.Server.User,
		Auth:            auth,
		HostKeyCallback: hostKeyCallbk,
	}
	addr = fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		log.Fatalln(err.Error())
	}
	sshClient = client
}
