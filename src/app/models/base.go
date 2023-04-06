package models

import (
	"app/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, "host=postgres user="+config.Config.DbUser+" dbname="+config.Config.DbName+" password="+config.Config.Password+" sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		uuid TEXT NOT NULL UNIQUE,
		name TEXT,
		email TEXT,
		password TEXT,
		created_at TIMESTAMP
	)`, tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		content TEXT,
		user_id INTEGER,
		created_at TIMESTAMP
	)`, tableNameTodo)

	_, err = Db.Exec(cmdT)
	if err != nil {
		log.Fatalln(err)
	}

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		uuid TEXT NOT NULL UNIQUE,
		email TEXT,
		user_id TEXT,
		created_at TIMESTAMP
	)`, tableNameSession)

	_, err = Db.Exec(cmdS)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
