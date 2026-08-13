package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB


type Bioskop struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

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
	router.POST("/bioskop", createBioskop)
	router.GET("/get-data-bioskop", getDataBioskop)
	router.PUT("/edit-data-bioskop/:id", updateDataBioskop)
	router.DELETE("/delete-data-bioskop/:id", deleteDataBioskop)

	// menjalankan server dan menentukan port nya
	router.Run(":8080")
}

func createBioskop(c *gin.Context) {

	var bioskop Bioskop

	// Mengambil data JSON
	err := c.ShouldBindJSON(&bioskop)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Format JSON tidak valid",
		})

		return
	}

	// memvalidasi nama bioskop
	if bioskop.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama tidak boleh kosong",
		})

		return
	}

	// memvalidasi lokasi
	if bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lokasi tidak boleh kosong",
		})

		return
	}

	// contoh untuk insert body
	// {
    // 	"nama": "XXI Pakuwon",
    // 	"lokasi": "Bekasi Barat",
    // 	"rating": 7
	// }

	// query insert ke database
	query := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err = db.QueryRow(
		query,
		bioskop.Nama,
		bioskop.Lokasi,
		bioskop.Rating,
	).Scan(&bioskop.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menyimpan data bioskop",
			"error":   err.Error(),
		})

		return
	}

	// Response yang di terima
	c.JSON(http.StatusCreated, gin.H{
		"message": "Bioskop berhasil ditambahkan",
		"data":    bioskop,
	})
}

func getDataBioskop(c *gin.Context) {
	var bioskop []Bioskop

	query := `
		SELECT id, nama, lokasi, rating
		FROM bioskop;
	`

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data bioskop",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var data Bioskop

		err := rows.Scan(
			&data.ID,
			&data.Nama,
			&data.Lokasi,
			&data.Rating,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal membaca data bioskop",
				"error":   err.Error(),
			})
			return
		}

		bioskop = append(bioskop, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data bioskop",
		"data":    bioskop,
	})
}

func updateDataBioskop(c *gin.Context) {
	var bioskop Bioskop

	id := c.Param("id")

	// Ambil JSON dari Postman
	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Format JSON tidak valid",
		})
		return
	}

	query := `
		UPDATE bioskop
		SET nama = $1,
		    lokasi = $2,
		    rating = $3
		WHERE id = $4
	`

	_, err := db.Exec(
		query,
		bioskop.Nama,
		bioskop.Lokasi,
		bioskop.Rating,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengubah data bioskop",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data bioskop berhasil diubah",
	})
}

func deleteDataBioskop(c *gin.Context) {
	id := c.Param("id")

	query := `
		DELETE FROM bioskop
		WHERE id = $1
	`

	result, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus data bioskop",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengecek data yang dihapus",
			"error":   err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data bioskop tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data bioskop berhasil dihapus",
	})
}