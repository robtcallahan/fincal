package repository

import (
	"vue-register/pkg/models"

	"gorm.io/gorm"
)

// QueryRepo represent the repositories
type QueryRepo interface {
	CreateDB(dbName string) (*gorm.DB, error)

	GetColumns() []models.Column
	GetCategoryColumns() []models.Column

	GetMerchants() []models.Merchant
	GetMerchantsAndColumns() []models.MerchantsAndColumns
	CreateMerchant(m *models.Merchant)
	UpdateMerchant(m *models.Merchant, column string, value string)
	DeleteMerchant(m *models.Merchant)

	GetTransactions() []models.Transaction
	SaveTransaction(trans *models.Transaction)
	UpdateTransactionTables(trans []*models.Transaction)

	GetLookupData() []*models.DataRow
	GetNameMapToColumn() map[string]string

	PrintData()
	PrintTable(table string)
}

func ColumnNames(cats []models.Column) *[]string {
	names := make([]string, 0, len(cats))
	for _, c := range cats {
		names = append(names, c.Name)
	}
	return &names
}
