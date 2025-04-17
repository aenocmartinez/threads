package database

import (
	_ "github.com/lib/pq"
)

// var (
// 	db   *sql.DB
// 	once sync.Once
// )

// func GetDBPostreSQL() *sql.DB {
// 	once.Do(func() {
// 		dsn := fmt.Sprintf(
// 			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
// 			os.Getenv("DB_USER"),
// 			os.Getenv("DB_PASSWORD"),
// 			os.Getenv("DB_HOST"),
// 			os.Getenv("DB_PORT"),
// 			os.Getenv("DB_NAME"),
// 		)

// 		var err error
// 		db, err = sql.Open("postgres", dsn)
// 		if err != nil {
// 			log.Fatalf("❌ Error al abrir la conexión: %v", err)
// 		}

// 		if err := db.Ping(); err != nil {
// 			log.Fatalf("❌ Error al hacer ping a la base de datos: %v", err)
// 		}

// 		log.Println("✅ Conexión a la base de datos PostgreSQL establecida.")
// 	})

// 	return db
// }
