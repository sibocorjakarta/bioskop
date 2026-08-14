package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBioskop(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
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

}
