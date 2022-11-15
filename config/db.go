package config

import (
	"database/sql"
	"time"

	_ "github.com/joho/godotenv"
)

func Connection() *sql.DB {
	host := DbHost()
	port := DbPort()
	name := DbName()
	username := DbUsername()
	password := DbPassword()

	DBURL := `` + username + `:` + password + `@tcp(` + host + `:` + port + `)/` + name + `?parseTime=true`
	db, err := sql.Open("mysql", DBURL)
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
