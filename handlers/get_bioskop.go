package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sibocorjakarta/bioskop/models"
)

func GetBioskop(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var bioskop []models.Bioskop

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
			var data models.Bioskop

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

}
