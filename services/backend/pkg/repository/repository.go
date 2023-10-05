package repository

import (
	"fincal/pkg/models"

	"gorm.io/gorm"
)

// QueryRepo represent the repositories
type QueryRepo interface {
	CreateDB(dbName string) (*gorm.DB, error)

	GetUserSecrets(userID int) (*models.Secrets, error)
	SaveSecrets(secrets *models.Secrets) (*models.Secrets, error)

	GetUserByUsername(username string) (*models.User, error)

	GetCategories() []models.Category
	GetCategoriesForSelect() []models.CategoryForSelect
	CreateCategory(category *models.Category) *models.Category
	UpdateCategory(category *models.Category, column string, value interface{})
	DeleteCategory(m *models.Category)

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
