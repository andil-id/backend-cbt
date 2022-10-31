package config

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/joho/godotenv"
)

func Connection() *sql.DB {
	DBURL := `` + os.Getenv("DB_USER") + `:` + os.Getenv("DB_PASS") + `@tcp(` + os.Getenv("DB_HOST") + `:` + os.Getenv("DB_PORT") + `)/` + os.Getenv("DB_NAME") + `?parseTime=true`
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
