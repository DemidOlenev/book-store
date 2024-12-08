package repo

import (
	"book-store/resources"
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

type connectionDB struct {
	conn *sql.DB
}

var (
	Connection *connectionDB
	once     sync.Once
)

func GetConnDBInstance() *connectionDB {
	once.Do(func() {
		conn, err := establish()
		if err != nil {
			fmt.Println("Unable to established connection to DB")
			log.Fatal(err)
		}
		Connection = &connectionDB{
			conn: conn,
		}
	})
	return Connection
}

func establish() (*sql.DB, error) {
	conn_settings := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		resources.DB_USER, resources.DB_PASS, resources.DB_NAME, resources.DB_SSLMODE)
	db, err := sql.Open(resources.DB_DRIVER, conn_settings)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to DB was established")
	return db, nil
}

func (db *connectionDB) Close() {
	fmt.Println("Connection to DB was closed")
	db.conn.Close()
}

func (db connectionDB) Ping() {
	fmt.Println("Ping to DB was started")
	db.conn.Ping()
	fmt.Println("Ping to DB was finished")

}
