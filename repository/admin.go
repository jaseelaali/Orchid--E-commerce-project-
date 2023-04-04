package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func SalesReport() ([]models.Order, error) {
	var salesData []models.Order

	salesReportQuery := `SELECT * FROM orders WHERE order_status = 'success';`

	if err := database.DB.Raw(salesReportQuery).Scan(&salesData).Error; err != nil {
		return []models.Order{}, err
	}
	return salesData, nil
}
