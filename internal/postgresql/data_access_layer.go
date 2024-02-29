package postgresql

import (
	"cdnetwork/pkg/log"
	"cdnetwork/pkg/util"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

// sshDriver is a custom driver that uses an SSH listener for tunneling

// sshDriver is a custom driver that uses an SSH listener for tunneling
type sshDriver struct {
	listener net.Listener
}

// Open implements the sql/driver.Driver interface
func (d *sshDriver) Open(dsn string) (driver.Conn, error) {
	conn, err := d.listener.Accept()
	if err != nil {
		return nil, err
	}
	return &sshConn{conn}, nil
}

// sshConn is a custom connection that wraps a net.Conn
type sshConn struct {
	conn net.Conn
}

// Prepare implements the driver.Conn interface
func (c *sshConn) Prepare(query string) (driver.Stmt, error) {
	panic("not implemented")
}

// Close implements the driver.Conn interface
func (c *sshConn) Close() error {
	return c.conn.Close()
}

// Begin implements the driver.Conn interface
func (c *sshConn) Begin() (driver.Tx, error) {
	panic("not implemented")
}

type DataAccess struct {
	Db        *sql.DB
	Listener  net.Listener
	SshClient *ssh.Client
}

func NewDataAccessLayer(config util.TgsConfig) (*DataAccess, error) {

	sshClient, err := connectSSH(config.SshTunnel)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Failed to connect to SSH server: %v", err))
	}

	remoteAddr := fmt.Sprintf("%s:%s", config.Postgresql.PgHost, config.Postgresql.PgPort)

	listener, err := sshClient.Listen("tcp", remoteAddr)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Failed to create db  port forwarder: %v", err))
	}

	fmt.Println("register start")
	sql.Register("postgres1", &sshDriver{listener})
	fmt.Println("register end")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=prefer",
		config.Postgresql.PgUsername,
		config.Postgresql.PgPassword,
		config.Postgresql.PgHost,
		config.Postgresql.PgDb,
	)

	// connStr := fmt.Sprintf("host=%s port=5432 user=%s "+
	// 	"password=%s dbname=%s sslmode=prefer",
	// 	config.Postgresql.PgHost,
	// 	config.Postgresql.PgUsername,
	// 	config.Postgresql.PgPassword,
	// 	config.Postgresql.PgDb)

	fmt.Println("conn:" + connStr)

	db, err := sql.Open("postgres1", connStr)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Failed to connect to database: %v", err))
		panic(err)
	}

	// fmt.Println("ping")

	// err = db.Ping()
	// if err != nil {
	// 	log.LogFatal(fmt.Sprintf("Failed to ping database: %v", err))
	// }

	// fmt.Println("pinged")

	return &DataAccess{
		Db:        db,
		Listener:  listener,
		SshClient: sshClient,
	}, nil
}
