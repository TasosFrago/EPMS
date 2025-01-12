package db_connection

import (
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

type viaSSHDialer struct {
	client *ssh.Client
}

func (self *viaSSHDialer) Dial(addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

type CredentialConfig struct {
	Usrname    string
	Passwd     string
	ServerHost string
	ServerPort string
	DBHost     string
	DBName     string
}

type DBConn struct {
	Conn    *sql.DB
	Cleanup func()
}

func ConnectDBoSSH(config CredentialConfig) (*DBConn, error) {
	sshConfig := &ssh.ClientConfig{
		User: config.Usrname,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Passwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH Server
	fmt.Println("Trying to connect with ssh...")
	sshcon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort), sshConfig)
	if err != nil {
		return nil, fmt.Errorf("SSHTunnelDB: %w", err)
	}
	fmt.Println("Established ssh")
	fmt.Println("Trying to connect to mysql over tcp...")
	// Now we register the ViaSSHDialer with the ssh connection as a parameter
	mysql.RegisterDial("mysql+tcp", (&viaSSHDialer{sshcon}).Dial)

	// And now we can use our new driver with the regular mysql connection string tunneled through the SSH connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s)/%s", config.Usrname, config.Passwd, config.DBHost, config.DBName))
	if err != nil {
		sshcon.Close()
		return nil, fmt.Errorf("SSHTunnelDB: %w", err)
	}

	db.SetMaxOpenConns(6)
	db.SetMaxIdleConns(6)
	db.SetConnMaxLifetime(5 * time.Minute)


	fmt.Printf("Successfully connected to the db!\n")

	return &DBConn{
		Conn: db,
		Cleanup: func() {
			db.Close()
			sshcon.Close()
		},
	}, nil
}
