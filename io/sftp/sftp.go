package sftp

import (
	"fmt"
	"net"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SftpClient sftp对象
type SftpClient struct {
	raw  *sftp.Client
	conf *SftpConfig
}

// SftpConfig sftp配置
type SftpConfig struct {
	User   string
	Passwd string
	Host   string
	Port   int
	Key    string
}

// NewSftpClient 新建sftp对象
func NewSftpClient(user, password, host, key string, port int) *SftpClient {
	return &SftpClient{
		conf: &SftpConfig{
			User:   user,
			Passwd: password,
			Host:   host,
			Port:   port,
			Key:    key,
		},
	}
}

// Connect 连接sftp服务器
func (cli *SftpClient) Connect() (err error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(cli.conf.Passwd))

	clientConfig = &ssh.ClientConfig{
		User: cli.conf.User,
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 30 * time.Second,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", cli.conf.Host, cli.conf.Port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return
	}

	cli.raw = sftpClient

	return
}

// Open 打开文件
func (cli *SftpClient) Open(path string) (*sftp.File, error) {
	return cli.raw.Open(path)
}

// Close 关闭连接
func (cli *SftpClient) Close() error {
	return cli.raw.Close()
}
