package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sibocorjakarta/bioskop/handlers"
)

var db *sql.DB

func main() {

	// cara Koneksi PostgreSQL
	var err error

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL tidak tersedia")
	}

	db, err = sql.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database berhasil terkoneksi")

	// cara Membuat router Gin
	router := gin.Default()

	// membuat Endpoint router POST
	router.POST("/bioskop", handlers.CreateBioskop(db))
	router.GET("/get-data-bioskop", handlers.GetBioskop(db))
	router.PUT("/edit-data-bioskop/:id", handlers.UpdateBioskop(db))
	router.DELETE("/delete-data-bioskop/:id", handlers.DeleteBioskop(db))

	// menjalankan server dan menentukan port nya
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}
