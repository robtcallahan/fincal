package repository

import (
	"vue-register/pkg/models"

	"gorm.io/gorm"
)

// QueryRepo represent the repositories
type QueryRepo interface {
	CreateDB(dbName string) (*gorm.DB, error)

	GetCategories() []models.Category
	GetCategoriesForSelect() []models.CategoryForSelect
	CreateCategory(category *models.Category) *models.Category
	UpdatesCategory(category *models.Category, values interface{})

	GetMerchants() []models.Merchant
	GetMerchantsAndCategories() []models.MerchantsAndCategories
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

func ColumnNames(cats []models.Category) *[]string {
	names := make([]string, 0, len(cats))
	for _, c := range cats {
		names = append(names, c.Name)
	}
	return &names
}
