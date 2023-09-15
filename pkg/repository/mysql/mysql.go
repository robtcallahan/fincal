package mysql

import (
	"fmt"
	"gorm.io/gorm/clause"
	"vue-register/pkg/models"
	repo "vue-register/pkg/repository"

	"gorm.io/gorm"
)

type mysqlQueryRepo struct {
	Conn *gorm.DB
}

// NewMySQLQueryRepo ...
func NewMySQLQueryRepo(conn *gorm.DB) repo.QueryRepo {
	return &mysqlQueryRepo{
		Conn: conn,
	}
}

func (r *mysqlQueryRepo) GetTransactions() []models.Transaction {
	var trans []models.Transaction
	r.Conn.Order("date").Find(&trans)
	return trans
}

func (r *mysqlQueryRepo) SaveTransaction(trans *models.Transaction) {
	r.Conn.Save(trans)
}

// UpdateTransactionTables ...
func (r *mysqlQueryRepo) UpdateTransactionTables(trans []*models.Transaction) {
	_ = r.Conn.AutoMigrate(&models.Transaction{})

	for _, t := range trans {
		result := r.Conn.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(t)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}

// CreateDB ...
func (r *mysqlQueryRepo) CreateDB(dbName string) (*gorm.DB, error) {
	db := r.Conn.Exec("CREATE DATABASE " + dbName)
	return db, db.Error
}

func (r *mysqlQueryRepo) GetCategories() []models.Category {
	var cols []models.Category
	r.Conn.Where("is_category = 1").Order("column_index").Find(&cols)
	return cols
}

func (r *mysqlQueryRepo) CreateCategory(category *models.Category) *models.Category {
	tx := r.Conn.Create(category)
	category.ID = tx.Statement.Model.(*models.Category).ID
	return category
}

func (r *mysqlQueryRepo) UpdatesCategory(category *models.Category, values interface{}) {
	r.Conn.Model(&category).Where("id = ?", category.ID).Updates(values)
}

func (r *mysqlQueryRepo) GetCategoriesForSelect() []models.CategoryForSelect {
	var cols []models.Category
	r.Conn.Where("is_category = 1").Order("name").Find(&cols)

	var catsSelect []models.CategoryForSelect
	for _, c := range cols {
		s := models.CategoryForSelect{
			ID:    c.ID,
			Value: c.ID,
			Text:  c.Name,
		}
		catsSelect = append(catsSelect, s)
	}
	return catsSelect
}

// GetMerchants ...
func (r *mysqlQueryRepo) GetMerchants() []models.Merchant {
	var merch []models.Merchant
	r.Conn.Order("name").Find(&merch)
	return merch
}

// CreateMerchant ...
func (r *mysqlQueryRepo) CreateMerchant(m *models.Merchant) {
	result := r.Conn.Create(&models.Merchant{
		Name:     m.Name,
		BankName: m.BankName,
		ColumnID: m.ColumnID,
	})
	if result.Error != nil {
		panic(result.Error)
	}
}

func (r *mysqlQueryRepo) UpdateMerchant(m *models.Merchant, column, value string) {
	r.Conn.Model(m).Update(column, value)
}

func (r *mysqlQueryRepo) DeleteMerchant(m *models.Merchant) {
	r.Conn.Delete(m)
}

func (r *mysqlQueryRepo) GetMerchantsAndCategories() []models.MerchantsAndCategories {
	var merch []models.MerchantsAndCategories
	r.Conn.Order("name").Find(&merch)
	return merch
}

// GetLookupData ...
func (r *mysqlQueryRepo) GetLookupData() []*models.DataRow {
	var merchants []models.MerchantsAndCategories

	r.Conn.Preload("Category").Find(&merchants)

	var data []*models.DataRow
	for _, m := range merchants {
		data = append(data, &models.DataRow{
			Name:          m.Name,
			BankName:      m.BankName,
			ColumnName:    m.Column.Name,
			ColumnIndex:   m.Column.ColumnIndex,
			Color:         m.Column.Color,
			IsCategory:    m.Column.IsCategory,
			TaxDeductible: m.TaxDeductible,
		})
	}
	return data
}

// GetNameMapToColumn creates a map lookup from trans name to budget category/column names
func (r *mysqlQueryRepo) GetNameMapToColumn() map[string]string {
	cols := r.GetLookupData()

	transNameToColName := make(map[string]string)
	for _, c := range cols {
		transNameToColName[c.Name] = c.ColumnName
	}
	return transNameToColName
}

// PrintData ...
func (r *mysqlQueryRepo) PrintData() {
	var merchants []models.MerchantsAndCategories
	r.Conn.Preload("Category").Find(&merchants)

	fmt.Printf("[Num] %-35s %-30s %-30s %-s\n", "Bank Name", "Name", "Category Name", "Category Index")
	for i, m := range merchants {
		fmt.Printf("[%3d] %-35s %-30s %-30s %2d\n", i+1, m.BankName, m.Name, m.Column.Name, m.Column.ColumnIndex)
	}
}

// PrintTable ...
func (r *mysqlQueryRepo) PrintTable(table string) {
	switch table {
	case "merchants":
		var merchants []models.MerchantsAndCategories
		result := r.Conn.Find(&merchants)
		fmt.Printf("%d rows found\n", result.RowsAffected)
		for _, m := range merchants {
			fmt.Printf("%d %s %s %s\n", m.ID, m.BankName, m.Name, m.Column.Name)
		}
	}
}
