package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sibocorjakarta/bioskop/models"
)

func CreateBioskop(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var bioskop models.Bioskop

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

}
