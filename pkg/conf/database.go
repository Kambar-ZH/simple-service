package conf

import "fmt"

type Database struct {
	Enable   bool
	Driver   string // mysql postgres
	Name     string // db name
	Host     string // ipaddr, hostname, or "0.0.0.0"
	HostName string // ipaddr, hostname, or "0.0.0.0"
	Port     int    // must be in range 1..65535
	User     string // db user_repo
	Password string // db password
	Dsn      string
}

func (db *Database) DSN() string {
	db.Dsn = fmt.Sprintf("%s://%s:%s@%s:%d/%s", db.Host, db.User, db.Password, db.HostName, db.Port, db.Name)
	return db.Dsn
}
