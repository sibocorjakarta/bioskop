package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sibocorjakarta/bioskop/handlers"
)

var db *sql.DB

func main() {

	// cara Koneksi PostgreSQL
	var err error
	connStr := "host=localhost port=5432 user=test password=password1234 dbname=bioskop_db sslmode=disable"

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	//  untuk Mengecek koneksi database
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	println("Database berhasil terhubung")

	// cara Membuat router Gin
	router := gin.Default()

	// membuat Endpoint router POST
	router.POST("/bioskop", handlers.CreateBioskop(db))
	router.GET("/get-data-bioskop", handlers.GetBioskop(db))
	router.PUT("/edit-data-bioskop/:id", handlers.UpdateBioskop(db))
	router.DELETE("/delete-data-bioskop/:id", handlers.DeleteBioskop(db))

	// menjalankan server dan menentukan port nya
	router.Run(":8080")
}
