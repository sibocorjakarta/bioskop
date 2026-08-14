package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sibocorjakarta/bioskop/models"
)

func UpdateBioskop(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var bioskop models.Bioskop

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

		result, err := db.Exec(
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
		rowsAffected, err := result.RowsAffected()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal mengecek data",
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
			"message": "Data bioskop berhasil diubah",
		})
	}

}
