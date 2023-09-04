package handler

import (
	"vue-register/pkg/driver"
	"vue-register/pkg/models"
	"vue-register/pkg/repository"
	"vue-register/pkg/repository/mysql"
	"vue-register/pkg/repository/postgres"
)

// Query ...
type Query struct {
	repo repository.QueryRepo
}

// NewQueryHandler ...
func NewQueryHandler(db *driver.DB) *Query {
	var repo repository.QueryRepo

	switch db.DBType {
	case driver.MySQL:
		repo = mysql.NewMySQLQueryRepo(db.SQL)
	case driver.PostgreSQL:
		repo = postgres.NewPostgreSQLQueryRepo(db.SQL)
	}
	return &Query{
		repo: repo,
	}
}

func (q *Query) GetTransactions() []models.Transaction {
	return q.repo.GetTransactions()
}

func (q *Query) SaveTransaction(trans *models.Transaction) {
	q.repo.SaveTransaction(trans)
}

func (q *Query) UpdateTransactionTables(trans []*models.Transaction) {
	q.repo.UpdateTransactionTables(trans)
}

// GetColumns get all columns
func (q *Query) GetColumns() []models.Column {
	return q.repo.GetColumns()
}

func (q *Query) GetCategoryColumns() []models.Column {
	return q.repo.GetCategoryColumns()
}

// GetMerchants get all merchants
func (q *Query) GetMerchants() []models.Merchant {
	return q.repo.GetMerchants()
}

func (q *Query) GetMerchantsAndColumns() []models.MerchantsAndColumns {
	return q.repo.GetMerchantsAndColumns()
}

// CreateMerchant ...
func (q *Query) CreateMerchant(m *models.Merchant) {
	q.repo.CreateMerchant(m)
}

func (q *Query) UpdateMerchant(m *models.Merchant, column, value string) {
	q.repo.UpdateMerchant(m, column, value)
}

func (q *Query) DeleteMerchant(m *models.Merchant) {
	q.repo.DeleteMerchant(m)
}

// GetLookupData ...
func (q *Query) GetLookupData() []*models.DataRow {
	return q.repo.GetLookupData()
}

// GetNameMapToColumn creates a map lookup from trans name to budget category/column names
func (q *Query) GetNameMapToColumn() map[string]string {
	return q.repo.GetNameMapToColumn()
}

// PrintData ...
func (q *Query) PrintData() {
	q.repo.PrintData()
}

// PrintTable ...
func (q *Query) PrintTable(table string) {
	q.repo.PrintTable(table)
}
