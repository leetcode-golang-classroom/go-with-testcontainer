package db_client

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Profile struct {
	Name string
}
type DbClient struct {
	ConnString string
	Db         *sql.DB
}

func NewDbClient(connString string) (*DbClient, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successsfult pinged the database")
	return &DbClient{
		ConnString: connString,
		Db:         db,
	}, nil
}
