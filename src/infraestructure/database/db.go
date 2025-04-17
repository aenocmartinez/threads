package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		fmt.Println(dsn)

		var err error
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("❌ Error al abrir la conexión MySQL: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("❌ Error al hacer ping a la base de datos MySQL: %v", err)
		}

		log.Println("✅ Conexión a la base de datos MySQL establecida.")
	})

	return db
}
