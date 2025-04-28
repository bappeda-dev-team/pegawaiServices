package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	log.Printf("Mencoba koneksi ke database %s di %s:%s...", dbName, dbHost, dbPort)

	if dbHost == "" {
		dbHost = "localhost"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error membuka koneksi database: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Gagal terhubung ke database dalam 10 detik: %v", err)
		log.Printf("Mencoba koneksi ulang...")

		ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err = db.PingContext(ctx)
		if err != nil {
			db.Close()
			log.Fatalf("Koneksi database gagal setelah percobaan ulang: %v", err)
		}
	}

	log.Printf("Berhasil terhubung ke database %s", dbName)
	log.Printf("Max Open Connections: %d", db.Stats().MaxOpenConnections)
	log.Printf("Open Connections: %d", db.Stats().OpenConnections)
	log.Printf("In Use Connections: %d", db.Stats().InUse)
	log.Printf("Idle Connections: %d", db.Stats().Idle)

	return db
}

func RunFlyway() {
	godotenv.Load()

	cmd := exec.Command("flyway", "-locations=filesystem:db/migrations", "migrate")

	newPath := "/opt/homebrew/opt/openjdk/bin:" + os.Getenv("PATH")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbHost == "" || dbPort == "" || dbName == "" || dbUser == "" {
		log.Fatal("Environment variables DB_HOST, DB_PORT, DB_NAME, DB_USER must be set")
	}

	cmd.Env = append(os.Environ(),
		"PATH="+newPath,
		"FLYWAY_URL=jdbc:postgresql://"+dbHost+":"+dbPort+"/"+dbName+"?sslmode=disable",
		"FLYWAY_USER="+dbUser,
		"FLYWAY_PASSWORD="+dbPassword,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Migrasi gagal dijalankan: %v", err)
	} else {
		log.Println("Migrasi berhasil dijalankan")
	}

}
